# Templ of Horse Lasers HTMX Example

- [HTMX](https://htmx.org/)
- [Hyperscript](https://hyperscript.org/)
- [Templ](https://templ.guide)
- [Tailwind](https://tailwindcss.com/)
- [GORM](https://gorm.io/)
- [SQLite](https://gorm.io/docs/connecting_to_the_database.html#SQLite)
- [Air](https://github.com/cosmtrek/air)

[HTMX](https://htmx.org/) and [Hyperscript](https://hyperscript.org/) libraries are stored locally in the `.dist` folder.

### Developer Experience Issues

- The VSCode extensions htmx-tags and \_hyperscript don't appear to work with templ files
- Changing templ file association to be seen as html breaks the templ-vscode extension integration
- For now, this project prefers to have templ-vscode working
- Follow [this github issue](https://github.com/a-h/templ/issues/407) for details from templ

### Frontend

When the server is running, the GO Echo endpoints will return the entire HTML document. HTMX attributes will request and update portions of the document. No separate frontend repo or client app is required.

### TEMPL: HTML Templating Engine

This example uses the [TEMPL templating engine](https://templ.guide/) to create and deliver HTML UX components to the requesting client.

### Tailwind Styles

This project uses [tailwind](https://tailwindcss.com/docs/installation) for styling shortcuts, and to compile the tailwind.css file for styles.

### Compiling & Hot Reloading with Air

This GO project uses [air](https://github.com/cosmtrek/air) for hot reloading. The .air.toml file is configured according to the [templ hot reloading instructions for air](https://templ.guide/commands-and-tools/hot-reload) with the `--proxy` flag to automatically reload pages when changes are detected. **HOWEVER** this does not currently work with air, and manual browser refresh is required after changes.

The `cmd` in the .air.toml handles generating the templ files and generating the tailwind.css file.

### Hyperscript - For Client Interactivity

- For interactivity that does not need to go through the server, this project uses [Hyperscript](https://hyperscript.org/).
- This is great for pure UX state (such as toggling ux element visibility) without needing a state management library

### Database - GORM + SQLite In Memory

- This project uses [GORM](https://gorm.io/) to interact with the database and comes with an in-memory implementation of [SQLite](https://gorm.io/docs/connecting_to_the_database.html#SQLite)
- The leaderboard example stores these values in the in-memory SQLite database

## Development

### VS Code:

See the `.vscode/settings.json` for optimizations that have been added for improved developer experience.

- `"editor.formatOnSave"` entry adds support for templ file formatting and tailwind integration
- ~~`"files.associations"` entry addition allows the vscode [htmx-tags](https://marketplace.visualstudio.com/items?itemName=otovo-oss.htmx-tags) extension to work on .templ files~~

### Runners:

Open terminal sessions for each of these hot reloading / watching processes to auto-restart the server on change, and to auto-rebuild the templ templates when they are updated.

```bash
# Install required node dependencies
$ npm i

# Start the hot reloading GO dev server
# The .air.toml includes commands to watch and generate templ and tailwind css files
$ air

# OPTIONAL - Run the invididual runners separately
# Tailwind compile and watch for changes
$ npx tailwindcss -o ./dist/tailwind.css --watch

# templ template changes / autogenerate and watch for changes
# the --proxy flag auto-reloads the browser
$ templ generate --proxy http://localhost:1323 --watch
```

## Production

```bash
# Optimize Tailwind styles with minification
$ npx tailwindcss -o ./dist/tailwind.css --minify
```
