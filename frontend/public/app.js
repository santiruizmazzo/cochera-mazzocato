import TenantsCollection from "./components/TenantsCollection.js";
import TenantMiniCard from "./components/TenantMiniCard.js";
import TenantFullCard from "./components/TenantFullCard.js";
import TenantForm from "./components/TenantForm.js";
import SlotsCollection from "./components/SlotsCollection.js";
import SlotCard from "./components/SlotCard.js";
import SlotForm from "./components/SlotForm.js";
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
    const link = event.target.closest("[data-link]");

    if (link) {
      event.preventDefault();
      router.navigateTo(link.href);
    }
  });

  document.body.addEventListener("navigate", (event) => {
    router.navigateTo(event.detail.href);
  });

  router.route();
});
