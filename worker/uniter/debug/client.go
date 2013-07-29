// Copyright 2013 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package debug

import (
	"strings"
)

// ClientScript returns a bash script suitable for executing
// on the unit system to intercept hooks via tmux shell.
func (c *DebugHooksContext) ClientScript(hooks []string) string {
	s := strings.Replace(debugHooksClientScript, "{unit_name}", c.Unit, -1)
	s = strings.Replace(s, "{entry_flock}", c.ClientFileLock(), -1)
	s = strings.Replace(s, "{exit_flock}", c.ClientExitFileLock(), -1)
	s = strings.Replace(s, "{hook_args}", strings.Join(hooks, " "), -1)
	return s
}

const debugHooksClientScript = `#!/bin/bash
(
# Lock the juju-<unit>-debug lockfile.
flock -n 8 || (echo "Failed to acquire {entry_flock}: unit is already being debugged" 2>&1; exit 1)
(
# Close the inherited lock FD, or tmux will keep it open.
exec 8>&-

# Write out the debug-hooks args.
echo "{hook_args}" > {entry_flock}

# Lock the juju-<unit>-debug-exit lockfile.
flock -n 9 || exit 1

# Wait for tmux to be installed.
while [ ! -f /usr/bin/tmux ]; do
    sleep 1
done

if [ ! -f ~/.tmux.conf ]; then
        if [ -f /usr/share/byobu/profiles/tmux ]; then
                # Use byobu/tmux profile for familiar keybindings and branding
                echo "source-file /usr/share/byobu/profiles/tmux" > ~/.tmux.conf
        else
                # Otherwise, use the legacy juju/tmux configuration
                cat > ~/.tmux.conf <<END

# Status bar
set-option -g status-bg black
set-option -g status-fg white

set-window-option -g window-status-current-bg red
set-window-option -g window-status-current-attr bright

set-option -g status-right ''

# Panes
set-option -g pane-border-fg white
set-option -g pane-active-border-fg white

# Monitor activity on windows
set-window-option -g monitor-activity on

# Screen bindings, since people are more familiar with that.
set-option -g prefix C-a
bind C-a last-window
bind a send-key C-a

bind | split-window -h
bind - split-window -v

# Fix CTRL-PGUP/PGDOWN for vim
set-window-option -g xterm-keys on

# Prevent ESC key from adding delay and breaking Vim's ESC > arrow key
set-option -s escape-time 0

END
        fi
fi

(
    # Close the inherited lock FD, or tmux will keep it open.
    exec 9>&-
    # The beauty below is a workaround for a bug in tmux (1.5 in Oneiric) or
    # epoll that doesn't support /dev/null or whatever.  Without it the
    # command hangs.
    tmux new-session -d -s {unit_name} 2>&1 | cat > /dev/null || true
    exec tmux attach -t {unit_name}
)
) 9>{exit_flock}
) 8>{entry_flock}
exit $?
`
