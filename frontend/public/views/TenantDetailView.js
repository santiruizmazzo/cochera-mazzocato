import AbstractView from "./AbstractView.js";
import TenantFullCard from "../components/TenantFullCard.js";
import ContentSection from "../components/ContentSection.js";
import TenantForm from "../components/TenantForm.js";

export default class extends AbstractView {
  constructor(params) {
    super(params);
  }

  renderWithin(mainElement) {
    this.fetchTenant();

    this.homeSection = new ContentSection();

    this.homeSection.title = "Info del inquilino";
    this.homeSection.buttonText = "Editar";
    this.homeSection.showGoBackIcon = true;

    this.tenantForm = new TenantForm();
    this.tenantForm.mode = "edit";
    this.homeSection.modalTitle = "Editar inquilino";
    this.homeSection.modalContent = this.tenantForm;

    this.tenantFullCard = new TenantFullCard();
    this.homeSection.content = this.tenantFullCard;

    mainElement.appendChild(this.homeSection);

    this.setUpEventListeners();
  }

  async fetchTenant() {
    const TENANT_URL = `${import.meta.env.VITE_API_URL}/api/tenants/${this.params.id}`;

    await fetch(TENANT_URL)
      .then((response) => {
        if (!response.ok) {
          throw new Error(`Error ${response.status}`);
        }
        return response.json();
      })
      .then((tenant) => {
        this.tenantFullCard.load(tenant);
        this.tenantForm.load(tenant);
      })
      .catch((error) => {
        this.homeSection.errorMessage = error;
      });
  }

  setUpEventListeners() {
    this.homeSection.addEventListener("tenants:updated", async (event) => {
      this.homeSection.modal.close();
      this.tenantFullCard.load(event.detail);
      this.tenantForm.load(event.detail);
    });
  }
}
