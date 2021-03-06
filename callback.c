#include "callback.h"

extern BOOL goFrameFunc();
extern BOOL goRenderFunc();
extern BOOL goFocusLostFunc();
extern BOOL goFocusGainFunc();
extern BOOL goGfxRestoreFunc();
extern BOOL goExitFunc();

void setFrameFunc(HGE_t *h, HGE_FuncState_t ff) {
	HGE_System_SetStateFunc(h, ff, goFrameFunc);
}

void setRenderFunc(HGE_t *h, HGE_FuncState_t ff) {
	HGE_System_SetStateFunc(h, ff, goRenderFunc);
}

void setFocusLostFunc(HGE_t *h, HGE_FuncState_t ff) {
	HGE_System_SetStateFunc(h, ff, goFocusLostFunc);
}

void setFocusGainFunc(HGE_t *h, HGE_FuncState_t ff) {
	HGE_System_SetStateFunc(h, ff, goFocusGainFunc);
}

void setGfxRestoreFunc(HGE_t *h, HGE_FuncState_t ff) {
	HGE_System_SetStateFunc(h, ff, goGfxRestoreFunc);
}

void setExitFunc(HGE_t *h, HGE_FuncState_t ff) {
	HGE_System_SetStateFunc(h, ff, goExitFunc);
}

