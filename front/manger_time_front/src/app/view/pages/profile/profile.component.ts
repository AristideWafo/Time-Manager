import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { ApiService } from '../../../services/api.service';
import { Router } from '@angular/router';

interface UserProfile {
  id?: number;
  firstName: string;
  lastName: string;
  email: string;
  phone?: string;
  position?: string;
  department?: string;
}

@Component({
  selector: 'app-profile',
  standalone: true,
  imports: [CommonModule, FormsModule],
  styleUrls: ['./profile.component.css'],
  template: `
  <div class="profile-container">
    <div class="profile-card">
      <div class="profile-header">
        <h1>Mon Profil</h1>
        <div class="profile-actions">
          <button 
            class="btn btn-secondary" 
            (click)="toggleEditMode()"
            [disabled]="isLoading">
            {{ isEditMode ? 'Annuler' : 'Modifier' }}
          </button>
          <button 
            class="btn btn-danger" 
            (click)="confirmDelete()"
            [disabled]="isLoading">
            Supprimer le compte
          </button>
        </div>
      </div>

      <div class="profile-content">
        <!-- Mode lecture -->
        <div *ngIf="!isEditMode" class="profile-view">
          <div class="profile-info">
            <div class="info-item">
              <label>Prénom :</label>
              <span>{{ userProfile.firstName || 'Non renseigné' }}</span>
            </div>
            <div class="info-item">
              <label>Nom :</label>
              <span>{{ userProfile.lastName || 'Non renseigné' }}</span>
            </div>
            <div class="info-item">
              <label>Email :</label>
              <span>{{ userProfile.email || 'Non renseigné' }}</span>
            </div>
            <div class="info-item">
              <label>Téléphone :</label>
              <span>{{ userProfile.phone || 'Non renseigné' }}</span>
            </div>
            <div class="info-item">
              <label>Poste :</label>
              <span>{{ userProfile.position || 'Non renseigné' }}</span>
            </div>
            <div class="info-item">
              <label>Département :</label>
              <span>{{ userProfile.department || 'Non renseigné' }}</span>
            </div>
          </div>
        </div>

        <!-- Mode édition -->
        <form *ngIf="isEditMode" class="profile-form" (ngSubmit)="updateProfile()">
          <div class="form-group">
            <label for="firstName">Prénom *</label>
            <input 
              type="text" 
              id="firstName" 
              [(ngModel)]="editProfile.firstName" 
              name="firstName"
              required>
          </div>

          <div class="form-group">
            <label for="lastName">Nom *</label>
            <input 
              type="text" 
              id="lastName" 
              [(ngModel)]="editProfile.lastName" 
              name="lastName"
              required>
          </div>

          <div class="form-group">
            <label for="email">Email *</label>
            <input 
              type="email" 
              id="email" 
              [(ngModel)]="editProfile.email" 
              name="email"
              required>
          </div>

          <div class="form-group">
            <label for="phone">Téléphone</label>
            <input 
              type="tel" 
              id="phone" 
              [(ngModel)]="editProfile.phone" 
              name="phone">
          </div>

          <div class="form-group">
            <label for="position">Poste</label>
            <input 
              type="text" 
              id="position" 
              [(ngModel)]="editProfile.position" 
              name="position">
          </div>

          <div class="form-group">
            <label for="department">Département</label>
            <input 
              type="text" 
              id="department" 
              [(ngModel)]="editProfile.department" 
              name="department">
          </div>

          <div class="form-actions">
            <button type="submit" class="btn btn-primary" [disabled]="isLoading">
              {{ isLoading ? 'Sauvegarde...' : 'Sauvegarder' }}
            </button>
            <button type="button" class="btn btn-secondary" (click)="toggleEditMode()" [disabled]="isLoading">
              Annuler
            </button>
          </div>
        </form>
      </div>

      <!-- Messages d'erreur/succès -->
      <div *ngIf="message" class="message" [ngClass]="messageType">
        {{ message }}
      </div>
    </div>

    <!-- Modal de confirmation de suppression -->
    <div *ngIf="showDeleteModal" class="modal-overlay" (click)="cancelDelete()">
      <div class="modal" (click)="$event.stopPropagation()">
        <h3>Confirmer la suppression</h3>
        <p>Êtes-vous sûr de vouloir supprimer votre compte ? Cette action est irréversible.</p>
        <div class="modal-actions">
          <button class="btn btn-danger" (click)="deleteProfile()" [disabled]="isLoading">
            {{ isLoading ? 'Suppression...' : 'Supprimer définitivement' }}
          </button>
          <button class="btn btn-secondary" (click)="cancelDelete()" [disabled]="isLoading">
            Annuler
          </button>
        </div>
      </div>
    </div>
  </div>
  `
})
export class ProfileComponent implements OnInit {
  userProfile: UserProfile = {
    firstName: '',
    lastName: '',
    email: '',
    phone: '',
    position: '',
    department: ''
  };

  editProfile: UserProfile = { ...this.userProfile };
  isEditMode = false;
  isLoading = false;
  message = '';
  messageType: 'success' | 'error' = 'success';
  showDeleteModal = false;

  constructor(
    private apiService: ApiService,
    private router: Router
  ) { }

  ngOnInit() {
    this.loadProfile();
  }

  loadProfile() {
    // Pour le moment, on simule des données
    // Quand le backend sera prêt, remplacer par :
    // this.apiService.getUserProfile(userId).subscribe(...)

    this.userProfile = {
      id: 1,
      firstName: 'Druss',
      lastName: 'Guts',
      email: 'Druss.guts@gmail.com',
      phone: '0123456789',
      position: 'Développeur',
      department: 'IT'
    };

    this.editProfile = { ...this.userProfile };
  }

  toggleEditMode() {
    this.isEditMode = !this.isEditMode;
    if (this.isEditMode) {
      this.editProfile = { ...this.userProfile };
    }
    this.clearMessage();
  }

  updateProfile() {
    if (!this.isValidForm()) {
      this.showMessage('Veuillez remplir tous les champs obligatoires', 'error');
      return;
    }

    this.isLoading = true;
    this.clearMessage();

    // Simulation d'appel API
    // Quand le backend sera prêt, remplacer par :
    // this.apiService.updateUserProfile(this.userProfile.id!, this.editProfile).subscribe({
    //   next: (response) => { ... },
    //   error: (error) => { ... }
    // });

    setTimeout(() => {
      this.userProfile = { ...this.editProfile };
      this.isEditMode = false;
      this.isLoading = false;
      this.showMessage('Profil mis à jour avec succès !', 'success');
    }, 1000);
  }

  confirmDelete() {
    this.showDeleteModal = true;
  }

  cancelDelete() {
    this.showDeleteModal = false;
  }

  deleteProfile() {
    this.isLoading = true;

    // Simulation d'appel API
    // Quand le backend sera prêt, remplacer par :
    // this.apiService.deleteUserProfile(this.userProfile.id!).subscribe({
    //   next: () => { ... },
    //   error: (error) => { ... }
    // });

    setTimeout(() => {
      this.isLoading = false;
      this.showDeleteModal = false;
      localStorage.removeItem('access_token');
      this.router.navigate(['/login']);
    }, 1000);
  }

  private isValidForm(): boolean {
    return !!(this.editProfile.firstName &&
      this.editProfile.lastName &&
      this.editProfile.email);
  }

  private showMessage(text: string, type: 'success' | 'error') {
    this.message = text;
    this.messageType = type;
    setTimeout(() => this.clearMessage(), 5000);
  }

  private clearMessage() {
    this.message = '';
  }
}


