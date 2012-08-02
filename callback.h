#ifndef HGE_GO_CALLBACK_H
#define HGE_GO_CALLBACK_H

#include "hge_c.h"

extern BOOL goFrameFunc();
extern BOOL goRenderFunc();
extern BOOL goFocusLostFunc();
extern BOOL goFocusGainFunc();
extern BOOL goGfxRestoreFunc();
extern BOOL goExitFunc();

void setFrameFunc(HGE_t* h, HGE_FuncState_t ff);
void setRenderFunc(HGE_t* h, HGE_FuncState_t ff);
void setFocusLostFunc(HGE_t* h, HGE_FuncState_t ff);
void setFocusGainFunc(HGE_t* h, HGE_FuncState_t ff);
void setGfxRestoreFunc(HGE_t* h, HGE_FuncState_t ff);
void setExitFunc(HGE_t* h, HGE_FuncState_t ff);

#endif
