import HomeView from "../views/HomeView.js";
import PaymentsView from "../views/PaymentsView.js";
import TenantsView from "../views/TenantsView.js";
import TenantDetailView from "../views/TenantDetailView.js";

export default class Router {
  routes = [
    { path: "/", view: HomeView },
    { path: "/pagos", view: PaymentsView },
    { path: "/inquilinos", view: TenantsView },
    { path: "/inquilinos/:id", view: TenantDetailView },
  ];

  navigateTo(relativePath) {
    if (this.currentRouteIsSameAs(relativePath)) return;

    history.pushState(null, null, relativePath);
    this.updateSelectedNavLinks(relativePath);
    this.renderCurrentView();
  }

  currentRouteIsSameAs(relativePath) {
    return relativePath.endsWith(location.pathname);
  }

  updateSelectedNavLinks(relativePath) {
    const navLinks = document.querySelectorAll(".nav-link");

    navLinks.forEach((navLink) => {
      navLink.className = relativePath.startsWith(navLink.href)
        ? "selected nav-link"
        : "nav-link";
    });
  }

  async renderCurrentView() {
    const view = this.createViewForCurrentRoute();
    const main = document.querySelector("main");

    main.innerHTML = "";
    await view.renderWithin(main);

    // main.innerHTML = "";
    // main.appendChild(await view.render());

    // main.innerHTML = await view.getHtml();
    // view.setUpJavascript();
  }

  createViewForCurrentRoute() {
    const matchingRouteInfo = this.findMatchingRoute();
    const viewParams = this.createViewParams(matchingRouteInfo);

    return new matchingRouteInfo.route.view(viewParams);
  }

  findMatchingRoute() {
    const potentialMatches = this.routes.map((route) => {
      return {
        route: route,
        result: location.pathname.match(this.pathToRegex(route.path)),
      };
    });

    let match = potentialMatches.find(
      (potentialMatch) => potentialMatch.result !== null,
    );

    return match
      ? match
      : { route: this.routes[0], result: [location.pathname] };
  }

  createViewParams(match) {
    const values = match.result.slice(1);
    const keys = Array.from(match.route.path.matchAll(/:(\w+)/g)).map(
      (result) => result[1],
    );

    if (values.length === 0 || keys.length === 0) return null;

    return Object.fromEntries(
      keys.map((key, i) => {
        return [key, values[i]];
      }),
    );
  }

  pathToRegex(path) {
    return new RegExp(
      "^" + path.replace(/\//g, "\\/").replace(/:\w+/g, "(.+)") + "$",
    );
  }
}
