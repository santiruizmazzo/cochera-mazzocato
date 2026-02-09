import AbstractView from "./AbstractView.js";
import ContentSection from "../components/ContentSection.js";

export default class extends AbstractView {
  constructor(params) {
    super(params);
  }

  renderWithin(mainElement) {
    // this.fetchPayments();

    this.homeSection = new ContentSection();

    this.homeSection.title = "Pagos recibidos";
    this.homeSection.buttonText = "Nuevo pago";
    this.homeSection.modalTitle = "Registrar pago";

    mainElement.appendChild(this.homeSection);
  }

  // async fetchPayments() {
  //   const PAYMENTS_URL = import.meta.env.VITE_API_URL + "/api/payments";

  //   await fetch(PAYMENTS_URL)
  //     .then((response) => {
  //       if (!response.ok) {
  //         throw new Error(`Error ${response.status}`);
  //       }
  //       return response.json();
  //     })
  //     .then(({ data }) => {
  //       this.paymentsCollection.load(data);
  //     })
  //     .catch((error) => {
  //       this.homeSection.errorMessage = error;
  //     });
  // }
}
