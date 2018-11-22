# generator

Code generator for `winapi` package

## Status

WIP.

## Example

```
$ go run tools/generator/*.go -s build/user32.h -t build/user32.h
```

> Note: `-t build/user32.h` above is just a dummy option, thus can be ignored

Outputs:
```
BOOL BringWindowToTop(HWND hWnd)
LRESULT CallWindowProcW(WNDPROC lpPrevWndFunc, HWND hWnd, UINT Msg, WPARAM wParam, LPARAM lParam)
BOOL CloseWindow(HWND hWnd)
HWND CreateWindowExW(DWORD dwExStyle, LPCWSTR lpClassName, LPCWSTR lpWindowName, DWORD dwStyle, int X, int Y, int nWidth, int nHeight, HWND hWndParent, HMENU hMenu, HINSTANCE hInstance, LPVOID lpParam)
```

Function signatures defined in `build/user32.h` are parsed and printed as-is(except for semicolons).
