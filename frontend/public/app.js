import { VersionBox } from "./components/VersionBox.js";
import { TenantsList } from "./components/TenantsList.js";
import { TenantCard } from "./components/TenantCard.js";
import HomeView from "./views/HomeView.js";
import PaymentsView from "./views/PaymentsView.js";
import TenantsView from "./views/TenantsView.js";

const navigateTo = (url) => {
  history.pushState(null, null, url);
  router();
};

const router = async () => {
  const routes = [
    { path: "/", view: HomeView },
    { path: "/pagos", view: PaymentsView },
    { path: "/inquilinos", view: TenantsView },
  ];

  const potentialMatches = routes.map((route) => {
    return {
      route: route,
      isMatch: location.pathname === route.path,
    };
  });

  let match = potentialMatches.find((potentialMatch) => potentialMatch.isMatch);

  if (!match) {
    match = {
      route: routes[0],
      isMatch: true,
    };
  }

  const view = new match.route.view();

  document.querySelector("main").innerHTML = await view.getHtml();
  view.setUpJavascript();
};

window.addEventListener("popstate", router);

document.addEventListener("DOMContentLoaded", () => {
  document.body.addEventListener("click", (event) => {
    if (event.target.matches("[data-link]")) {
      event.preventDefault();
      navigateTo(event.target.href);
    }
  });

  router();
});
