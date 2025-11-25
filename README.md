# Confy

A simple tool for moving and editing configuration files easily

## Installation

confy can either be installed through the following package managers or by downloading the package from the releases tab and adding the package to path (or you can always clone and build by yourself)

> [!WARNING] 
> 
> confy isn't on the AUR yet and this part of the documentation is just filler

For Arch use AUR:

```bash
  yay -S confy-tool
```

## Features

- Backup old configs in case you change your mind
- Move configuration folders/files with just one command
- Track configuration history

## Documentation

Confy uses a module/config system where each file is treated as a config and the folder they are kept in is called a module eg in .config/waybar 'waybar' is the module and the 'config.jsonc' file contained in it is treated as a config.

Link to the full [Documentation](https://lordryns.github.io)



## Quick Reference

The command below sets replaces the waybar in the global config path set by the user to the new_waybar (gets renamed to waybar, only contents change) and the old gets backed up

```bash
confy set waybar /path/to/new_waybar
```

To rollback to the previous config all you have to do is use the rollback command

```bash
confy rollback waybar
```

This is just the tip of the iceberg of what confy can really do, here's a link to the full [Documentation](https://lordryns.github.io)

## Contributing

Contributions are always welcome, Just send a pull request!

## Appendix

Thanks for checking out this project, please leave if a star if you can spare the time, Reach me on [X](https://x.com/lordryns) or [Discord](https://discord.com/users/1015382973052358657) (username is @lordryns on both)