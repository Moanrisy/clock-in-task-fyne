# Clock in Task Fyne

Simple program to show current task that I'm working on

# How it work

## Fyne toolkit

Using fyne toolkit to create desktop app using go language. The app will show current task that I currently clocked-in using emacs

## Emacs function

```lisp
;; Hook for clock-in
(defun write-clock-in-title-hook()
"Write clock in title into a file"
(message (symbol-value 'org-clock-heading))
(append-to-file (symbol-value 'org-clock-heading) nil "~/clock-in-title")
(append-to-file "\n" nil "~/clock-in-title")
)

(add-hook 'org-clock-in-hook 'write-clock-in-title-hook)
```

## Power toys

This 3rd party app used to make window in window 11 can be used as 'always on top'
