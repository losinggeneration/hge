package hge

import (
	"strings"
)

const (
	strHEADERTAG   = "[HGESTRINGTABLE]"
	strFORMATERROR = "String table %s has incorrect format."
)

type StringTable struct {
	stringsMap map[string]string
}

func NewStringTable(filename string) *StringTable {
	st := new(StringTable)
	hge := Create(VERSION)

	st.stringsMap = make(map[string]string)

	f := hge.ResourceLoadString(filename)

	if f == nil || !strings.HasPrefix(*f, strHEADERTAG) {
		hge.System_Log(strFORMATERROR, filename)
		return nil
	}

	var (
		inComment, inIdentifier, inValue bool
		identifier, value                string
	)

	reader := strings.NewReader(*f)
	_, e := reader.Seek(int64(len(strHEADERTAG)), 0)

	if e != nil {
		hge.System_Log("Unable to seek past header tag")
		return nil
	}

	for b, e := reader.ReadByte(); e == nil; b, e = reader.ReadByte() {
		// we ignore whitespace
		if b == '\n' {
			inComment = false
			inIdentifier = false
			continue
		}

		// just continue if we're in a comment
		if inComment {
			continue
		}

		if inIdentifier {
			// break from the identifier when we get to whitespace
			if b == ' ' {
				inIdentifier = false
				continue
			}
			identifier += string(b)
			continue
		}

		if inValue {
			// We've found a backslash, figure out what to do with it
			if b == '\\' {
				// We need the next byte
				b, e = reader.ReadByte()
				if e != nil {
					// but break on an error
					break
				}

				switch b {
				case 'n':
					// insert a literal \n as the value
					value += "\\n"
				case '"':
					// insert a literal " as the value
					value += "\""
				case '\\':
					// insert a literal \ in the value
					value += "\\"
				default:
					// we don't have a special backslash, so just put a
					// backslash and the byte
					value += "\\" + string(b)
				}
				continue
			}

			// We've found the end of the value
			if b == '"' {
				// So assign it to the stringsMap
				st.stringsMap[identifier] = value
				// clear the identifier and value
				identifier, value = "", ""
				inValue = false
				continue
			}
			value += string(b)
			continue
		}

		// We found a comment
		if b == ';' {
			inComment = true
			continue
		}

		// ignore whitespace around identifiers and values
		if b == ' ' {
			continue
		}

		// start of an identifier (an identifier starts with a letter and has
		// any other character in it's name except whitespace)
		if (b > 'a' && b < 'z') || (b > 'A' && b < 'Z') {
			identifier = string(b)
			inIdentifier = true
		}

		// We search until we find the start of the value
		if b == '=' {
			for b, e := reader.ReadByte(); e == nil; b, e = reader.ReadByte() {
				// This means we've found the start of the value
				if b == '"' {
					inValue = true
					break
				}
			}
		}
	}

	return st
}

func (st StringTable) GetString(name string) string {
	return st.stringsMap[name]
}
