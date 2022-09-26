# go-spacer
Awesome-like workspace handeling for multimonitor systems written in go

## Installation

Download binaries, and put them into your path

Run go-spacer -index 1

Update .config/i3/config to have

`include output/*`

## Usage
Default configuration for controls is:

```
bindsym $mod+1 exec --no-startup-id "go-spacer -index 1"
...
bindsym $mod+0 exec --no-startup-id "go-spacer -index 10"

bindsym $mod+Shift+1 exec --no-startup-id "go-spacer -shift -index 1"
...
bindsym $mod+Shift+0 exec --no-startup-id "go-spacer -shift -index 10"

bindsym $mod+Control+1 exec --no-startup-id "go-spacer -shift -go -index 1"
...
bindsym $mod+Control+0 exec --no-startup-id "go-spacer -shift -go -index 10"
```

and for each output

```
workspace "{{add .index  1}}: 1" output {{.name}}
workspace "{{add .index  2}}: 2" output {{.name}}
workspace "{{add .index  3}}: 3" output {{.name}}
workspace "{{add .index  4}}: 4" output {{.name}}
workspace "{{add .index  5}}: 5" output {{.name}}
workspace "{{add .index  6}}: 6" output {{.name}}
workspace "{{add .index  7}}: 7" output {{.name}}
workspace "{{add .index  8}}: 8" output {{.name}}
workspace "{{add .index  9}}: 9" output {{.name}}
workspace "{{add .index  10}}: 10" output {{.name}}
```


Enjoy


## License

MIT


