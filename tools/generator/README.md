# generator

Code generator for `winapi` package

## Status

WIP.

## Example

```
$ go run tools/generator/*.go -d build/user32.h
```

Outputs:
```
type: BOOL -> int32
type: RECT -> struct {
        INT left
        INT top
        INT right
        INT bottom
}
function: BOOL BringWindowToTop(HWND hWnd)
function: LRESULT CallWindowProcW(WNDPROC lpPrevWndFunc, HWND hWnd, UINT Msg, WPARAM wParam, LPARAM lParam)
function: BOOL CloseWindow(HWND hWnd)
function: HWND CreateWindowExW(DWORD dwExStyle, LPCWSTR lpClassName, LPCWSTR lpWindowName, DWORD dwStyle, int X, int Y, int nWidth, int nHeight, HWND hWndParent, HMENU hMenu, HINSTANCE hInstance, LPVOID lpParam)
```

Function signatures and type definitions defined in `build/user32.h` are parsed and printed.
