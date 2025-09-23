import VersionBox from "./components/VersionBox.js";
import TenantsCollection from "./components/TenantsCollection.js";
import TenantCard from "./components/TenantCard.js";
import TenantForm from "./components/TenantForm.js";
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
