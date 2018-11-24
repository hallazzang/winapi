// copied from MSDN

typedef int32 BOOL;
typedef struct {
    INT left;
    INT top;
    INT right;
    INT bottom;
} RECT;

BOOL BringWindowToTop(
  HWND hWnd
);

LRESULT CallWindowProcW(
  WNDPROC lpPrevWndFunc,
  HWND    hWnd,
  UINT    Msg,
  WPARAM  wParam,
  LPARAM  lParam
);

BOOL CloseWindow(
  HWND hWnd
);

HWND CreateWindowExW(
  DWORD     dwExStyle,
  LPCWSTR   lpClassName,
  LPCWSTR   lpWindowName,
  DWORD     dwStyle,
  int       X,
  int       Y,
  int       nWidth,
  int       nHeight,
  HWND      hWndParent,
  HMENU     hMenu,
  HINSTANCE hInstance,
  LPVOID    lpParam
);
