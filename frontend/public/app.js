import TenantsCollection from "./components/TenantsCollection.js";
import TenantCard from "./components/TenantCard.js";
import TenantForm from "./components/TenantForm.js";
import ActivatableButton from "./components/ActivatableButton.js";
import ErrorBox from "./components/ErrorBox.js";
import iconsSprite from "./assets/icons.svg?raw";
import Router from "./routing/Router.js";

const router = new Router();

window.addEventListener("popstate", () => router.route());

document.addEventListener("DOMContentLoaded", () => {
  document.body.insertAdjacentHTML(
    "afterbegin",
    `<div style="display: none;">${iconsSprite}</div>`,
  );

  document.body.addEventListener("click", (event) => {
    if (event.target.matches("[data-link]")) {
      event.preventDefault();
      router.navigateTo(event.target.href);
    }
  });

  document.body.addEventListener("navigate", (event) => {
    router.navigateTo(event.detail.href);
  });

  router.route();
});
