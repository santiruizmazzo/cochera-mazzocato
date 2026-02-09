import AbstractView from "./AbstractView.js";
import ContentSection from "../components/ContentSection.js";
import TenantsCollection from "../components/TenantsCollection.js";
import TenantForm from "../components/TenantForm.js";

export default class extends AbstractView {
  constructor(params) {
    super(params);
  }

  renderWithin(mainElement) {
    this.fetchTenants();

    this.homeSection = new ContentSection();

    this.homeSection.title = "Inquilinos registrados";
    this.homeSection.buttonText = "Nuevo inquilino";

    this.tenantForm = new TenantForm();
    this.homeSection.modalTitle = "Registrar inquilino";
    this.homeSection.modalContent = this.tenantForm;

    this.tenantsCollection = new TenantsCollection();

    this.homeSection.content = this.tenantsCollection;

    mainElement.appendChild(this.homeSection);

    this.setUpEventListeners();
  }

  async fetchTenants() {
    const TENANTS_URL = import.meta.env.VITE_API_URL + "/api/tenants";

    await fetch(TENANTS_URL)
      .then((response) => {
        if (!response.ok) {
          throw new Error(`Error ${response.status}`);
        }
        return response.json();
      })
      .then(({ data }) => {
        this.tenantsCollection.load(data);
      })
      .catch((error) => {
        this.homeSection.errorMessage = error;
      });
  }

  setUpEventListeners() {
    this.homeSection.addEventListener("tenants:created", () => {
      this.homeSection.modal.close();
      this.fetchTenants();
    });
  }
}
