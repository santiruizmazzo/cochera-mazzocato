import HomeView from "../views/HomeView.js";
import PaymentsView from "../views/PaymentsView.js";
import TenantsView from "../views/TenantsView.js";

export default class Router {
  routes = [
    { path: "/", view: HomeView },
    { path: "/pagos", view: PaymentsView },
    { path: "/inquilinos", view: TenantsView },
  ];

  navigateTo(url) {
    if (url.endsWith(location.pathname)) return;

    history.pushState(null, null, url);
    this.route();
  }

  async route() {
    const potentialMatches = this.routes.map((route) => {
      return {
        route: route,
        isMatch: location.pathname === route.path,
      };
    });

    let match = potentialMatches.find(
      (potentialMatch) => potentialMatch.isMatch,
    );

    if (!match) {
      match = {
        route: routes[0],
        isMatch: true,
      };
    }

    const view = new match.route.view();

    document.querySelector("main").innerHTML = await view.getHtml();
    view.setUpJavascript();
  }
}
