import AbstractView from "./AbstractView.js";

export default class extends AbstractView {
  constructor() {
    super();
    this.setTitle("Inquilinos");
  }

  async getHtml() {
    return /*html*/ `
      <button class="open-button">Nuevo inquilino</button>
      <dialog class="modal">
        <form class="form" method="dialog">
          <h3>Crear nuevo inquilino</h3>
          <label>Nombre <input type="text"></label>
          <label>Email <input type="email"></label>
          <button type="submit">Crear</button>
        </form>
        <button class="close-button">Cerrar</button>
      </dialog>
      <h2 class="section-title">Inquilinos registrados</h2>
      <tenants-list></tenants-list>
    `;
  }

  setUpJavascript() {
    const modal = document.querySelector(".modal");
    const openModal = document.querySelector(".open-button");
    const closeModal = document.querySelector(".close-button");

    openModal.addEventListener("click", () => {
      modal.showModal();
    });

    closeModal.addEventListener("click", () => {
      modal.close();
    });
  }
}
