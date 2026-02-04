import AbstractView from "./AbstractView.js";

export default class extends AbstractView {
  constructor(params) {
    super(params);
  }

  async getHtml() {
    return /*html*/ `
      <section class="payments-view">
        <dialog id="new-payment-modal" closedby="any">
          <header>
            <h3>Modal generica</h3>
            <close-modal-button></close-modal-button>
          </header>
        </dialog>
      
        <header class="section-header">
          <h2>Pagos recibidos</h2>
          <open-modal-button></open-modal-button>
          <!-- <button class="new-action-btn open-modal-btn">
            <svg><use href="#person_edit"/></svg>
            Nuevo pago
          </button> -->
        </header>

        <div>Acá irían los pagos realizados</div>
      </section>
    `;
  }

  setUpJavascript() {
    const view = document.querySelector(".payments-view");
    const modal = document.querySelector("#new-payment-modal");

    view.addEventListener("open-modal", () => modal.showModal());
    view.addEventListener("close-modal", () => modal.close());
  }
}
