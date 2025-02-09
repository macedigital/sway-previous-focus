# sway-previous-focus

Static binary that marks previous window to quickly toggle between last recently used ones.

This project uses Sway IPC via [go-sway][1] library.

## Getting started

1. Compile source or put release binary into a folder in `$PATH` (e.g. `~/.local/bin`).
2. Start script while running in a Sway session (or enable systemd unit for auto re-starts).
3. Bind key in Sway configuration to jump to window marked with `_prev`.

## Credits

Inspired by <https://gitlab.com/wef/dotfiles/-/blob/c6cd8c0a7e003a2fa2c5f4188e23aad1cc90ce85/bin/i3-track-window-events> and
discussion at <https://www.reddit.com/r/swaywm/comments/u9txcn/does_anyone_have_interesting_uses_for_sways_marks/>.

[1]: https://github.com/joshuarubin/go-sway
