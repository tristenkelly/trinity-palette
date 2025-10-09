# The Trinity Palette
E-commerce platform/blog, storefront for art.

## Features

- User signup, login, email/password change  
- Token-based authentication (refresh, verify)  
- CRUD operations for “items” / “posts”  
- Blog page rendering  
- Shop handler / item listing  
- Serving static assets & templates  
- SQL database interactions via generated code  
- Infrastructure / deployment scaffolding  

## Tech Stack

- **Language**: Go  
- **Templating / HTTP**: Go’s `net/http` + HTML templates  
- **Database / SQL**: Using SQL files + code generation (via `sqlc`)  
- **Infrastructure**: Terraform / HCL (in `infra` folder)  
- **CI / Workflows**: GitHub Actions  
- **Security / Linting**: gosec setup (via `gosec.config.json`)  

## How It’s Built

1. **Routing & HTTP Handlers**  
   Each feature (login, signup, blog pages, shop, etc.) has its own handler file (e.g. `login_handler.go`, `shop_handler.go`) which handles HTTP requests, parses inputs, interacts with services, and writes output (HTML or JSON).

2. **Templates & Static Files**  
   - HTML templates are stored in `templates/`  
   - CSS, JS, image assets in `static/`  
   - Handlers render templates by passing data structs  

3. **Data Layer / Persistence**  
   - SQL schema / queries stored in `sql/`  
   - `sqlc.yaml` configures code generation  
   - Generated Go code in `internal` (or wherever configured) to provide typed DB access  

4. **Auth / Tokens**  
   - JWT or token based access (verify, refresh)  
   - Token endpoints: `verify_token.go`, `refresh_token.go`  
   - User identity and session logic encapsulated in handlers  

5. **Infrastructure & Deployment**  
   - `infra/` houses Terraform / HCL code for provisioning backend infrastructure  
   - GitHub Actions workflows (in `.github/workflows/`) automate builds, tests, deploys  

6. **Security / Audits**  
   - `gosec.config.json` configures the gosec static analyzer  
   - Linting, static security checks are part of the CI pipeline  