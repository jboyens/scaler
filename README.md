# Scaler

Scaler is a simple app that I wrote that scans the XDG_CONFIG_HOME looking for
`.tmpl` files. It then uses a Go template to do a replacement based on a YAML
config file.

Given a template like this (for the kitty terminal) at `$XDG_CONFIG_HOME/kitty/kitty.conf.tmpl`:

```text
font_family Iosevka Term SS05 Light
bold_font Iosevka Term SS05 Bold
italic_font Iosevka Term SS05 Light Italic
bold_italic_font Iosevka Term SS05 Bold Italic

font_size {{.font_size}}
```

and a YAML file like this at `$XDG_CONFIG_HOME/.scaler/config.yml`:

```yaml
---
apps:
  - name: kitty
    up:
      font_size: 15
    down:
      font_size: 10
```

running `scaler UP` will produce an `$XDG_CONFIG_HOME/kitty/kitty.conf` like this:

```text
font_family Iosevka Term SS05 Light
bold_font Iosevka Term SS05 Bold
italic_font Iosevka Term SS05 Light Italic
bold_italic_font Iosevka Term SS05 Bold Italic

font_size 15
```

running `scaler UP` will produce this:

```text
font_family Iosevka Term SS05 Light
bold_font Iosevka Term SS05 Bold
italic_font Iosevka Term SS05 Light Italic
bold_italic_font Iosevka Term SS05 Bold Italic

font_size 10
```

It's all absurdly simple, but I needed to change a ton of files for different
DPIs. I use [Kanshi](https://github.com/emersion/kanshi) to run the command
after changes to the outputs
