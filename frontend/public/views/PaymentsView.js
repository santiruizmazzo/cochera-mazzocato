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
    this.homeSection.buttonText = "Nuevo pago";
    this.homeSection.modalTitle = "Registrar pago";

    mainElement.appendChild(this.homeSection);
  }
}
