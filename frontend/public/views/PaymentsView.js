import AbstractView from "./AbstractView.js";
import ContentSection from "../components/ContentSection.js";

export default class extends AbstractView {
  constructor(params) {
    super(params);
  }

  renderWithin(mainElement) {
    // this.fetchPayments().then((payments) => {
    //   this.paymentsCollection.loadPayments(payments);
    // });

    this.homeSection = new ContentSection();

    this.homeSection.title = "Pagos recibidos";
    
    mainElement.appendChild(this.homeSection);
  }

  async getHtml() {
    return /*html*/ `
      <section class="payments-view">
        <custom-modal title="Registrar pago"></custom-modal>
      
        <header class="section-header">
          <h2>Pagos recibidos</h2>
          <open-modal-button>
            <svg><use href="#person_add"></svg>
            Nuevo pago
          </open-modal-button>
        </header>

        <div>Acá irían los pagos realizados</div>
        <div></div>
      </section>
    `;
  }
}
