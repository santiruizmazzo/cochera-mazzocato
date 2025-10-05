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

  pathToRegex(path) {
    return new RegExp(
      "^" + path.replace(/\//g, "\\/").replace(/:\w+/g, "(.+)") + "$",
    );
  }

  getParams(match) {
    const values = match.result.slice(1);
    const keys = Array.from(match.route.path.matchAll(/:(\w+)/g)).map(
      (result) => result[1],
    );

    return Object.fromEntries(
      keys.map((key, i) => {
        return [key, values[i]];
      }),
    );
  }

  navigateTo(url) {
    if (url.endsWith(location.pathname)) return;

    history.pushState(null, null, url);
    this.route();
  }

  async route() {
    const potentialMatches = this.routes.map((route) => {
      return {
        route: route,
        result: location.pathname.match(this.pathToRegex(route.path)),
      };
    });

    let match = potentialMatches.find(
      (potentialMatch) => potentialMatch.result !== null,
    );

    if (!match) {
      match = {
        route: this.routes[0],
        result: [location.pathname],
      };
    }

    const view = new match.route.view(this.getParams(match));

    document.querySelector("main").innerHTML = await view.getHtml();
    view.setUpJavascript();
  }
}
