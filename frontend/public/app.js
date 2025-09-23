import VersionBox from "./components/VersionBox.js";
import TenantsList from "./components/TenantsList.js";
import TenantCard from "./components/TenantCard.js";
import Router from "./routing/Router.js";

const router = new Router();

window.addEventListener("popstate", router.route);

document.addEventListener("DOMContentLoaded", () => {
  document.body.addEventListener("click", (event) => {
    if (event.target.matches("[data-link]")) {
      event.preventDefault();
      router.navigateTo(event.target.href);
    }
  });

  router.route();
});
