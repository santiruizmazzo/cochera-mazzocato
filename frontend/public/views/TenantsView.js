import AbstractView from "./AbstractView.js";

export default class extends AbstractView {
  constructor() {
    super();
  }

  getHtml() {
    return /*html*/ `
      <dialog id="tenant-creation-modal" class="modal" closedby="any">
        <h3>Registrar inquilino</h3>
        <button class="close-modal">
          <svg><use href="#close"/></svg>
        </button>
        <tenant-form></tenant-form>
      </dialog>

      <header class="tenants-header">
        <h2 class="section-title">Inquilinos registrados</h2>
        <button class="open-modal">
          <svg><use href="#person_add"/></svg>
          <p>Nuevo inquilino</p>
        </button>
      </header>
      
      <tenants-collection>
        <svg class="spinner"><use href="#spinner"/></svg>
      </tenants-collection>
    `;
  }

  setUpJavascript() {
    const modal = document.querySelector("#tenant-creation-modal");
    const openModal = document.querySelector(".open-modal");
    const closeModal = document.querySelector(".close-modal");
    const form = document.querySelector("tenant-form");

    openModal.addEventListener("click", () => modal.showModal());
    closeModal.addEventListener("click", () => {
      modal.close();
      form.clear();
    });
    modal.addEventListener("tenants:update", () => modal.close());
  }
}
