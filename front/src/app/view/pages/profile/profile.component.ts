import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { ApiService } from '../../../services/api.service';
import { Router } from '@angular/router';
import * as jwtDecodeModule from 'jwt-decode';

interface UserProfile {
  id?: string;
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
              [disabled]="isLoading"
            >
              {{ isEditMode ? 'Annuler' : 'Modifier' }}
            </button>

            <button
              class="btn btn-danger"
              (click)="confirmDelete()"
              [disabled]="isLoading"
            >
              Supprimer le compte
            </button>
          </div>
        </div>

        <div class="profile-content">
          <!-- Vue profil -->
            <div *ngIf="!isEditMode" class="profile-view">
            <div class="profile-info">
              <div>
                <label>Prénom :</label>
                <span>{{ userProfile.firstName || 'Non renseigné' }}</span>
              </div>

              <div>
                <label>Nom :</label>
                <span>{{ userProfile.lastName || 'Non renseigné' }}</span>
              </div>

              <div>
                <label>Email :</label>
                <span>{{ userProfile.email || 'Non renseigné' }}</span>
              </div>

              <div>
                <label>Téléphone :</label>
                <span>{{ userProfile.phone || 'Non renseigné' }}</span>
              </div>

              <div>
                <label>Poste :</label>
                <span>{{ userProfile.position || 'Non renseigné' }}</span>
              </div>

              <div>
                <label>Département :</label>
                <span>{{ userProfile.department || 'Non renseigné' }}</span>
              </div>
            </div>
          </div>

          <!-- Formulaire d’édition -->
          <form *ngIf="isEditMode" (ngSubmit)="updateProfile()">
            <div>
              <label>Prénom *</label>
              <input [(ngModel)]="editProfile.firstName" name="firstName" required />
            </div>

            <div>
              <label>Nom *</label>
              <input [(ngModel)]="editProfile.lastName" name="lastName" required />
            </div>

            <div>
              <label>Email *</label>
              <input [(ngModel)]="editProfile.email" name="email" required />
            </div>

            <div>
              <label>Téléphone</label>
              <input [(ngModel)]="editProfile.phone" name="phone" />
            </div>

            <div>
              <label>Poste</label>
              <input [(ngModel)]="editProfile.position" name="position" />
            </div>

            <div>
              <label>Département</label>
              <input [(ngModel)]="editProfile.department" name="department" />
            </div>

            <button type="submit" [disabled]="isLoading">
              {{ isLoading ? 'Sauvegarde...' : 'Sauvegarder' }}
            </button>

            <button
              type="button"
              (click)="toggleEditMode()"
              [disabled]="isLoading"
            >
              Annuler
            </button>
          </form>

          <div *ngIf="message" [ngClass]="messageType">
            {{ message }}
          </div>
        </div>

        <!-- Modal de suppression -->
        <div *ngIf="showDeleteModal" class="modal-overlay" (click)="cancelDelete()">
          <div class="modal" (click)="$event.stopPropagation()">
            <h3>Confirmer la suppression</h3>
            <p>Êtes-vous sûr de vouloir supprimer votre compte ?</p>

            <button (click)="deleteProfile()" [disabled]="isLoading">
              {{ isLoading ? 'Suppression...' : 'Supprimer définitivement' }}
            </button>

            <button (click)="cancelDelete()" [disabled]="isLoading">
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
  userId!: string;

  constructor(
    private apiService: ApiService,
    private router: Router
  ) { }

  ngOnInit() {
    const token = localStorage.getItem('access_token');
    if (!token) {
      this.router.navigate(['/login']);
      return;
    }

    try {
      const decoded: any = (jwtDecodeModule as any).default(token);
      this.userId = decoded._id;
      this.loadProfile();
    } catch {
      localStorage.removeItem('access_token');
      this.router.navigate(['/login']);
    }
  }

  loadProfile() {
    this.apiService.getUser(this.userId).subscribe({
      next: (res: any) => {
        this.userProfile = res;
        this.editProfile = { ...res };
      },
      error: () => this.showMessage('Impossible de charger le profil.', 'error')
    });
  }

  toggleEditMode() {
    this.isEditMode = !this.isEditMode;
    if (this.isEditMode) this.editProfile = { ...this.userProfile };
    this.clearMessage();
  }

  updateProfile() {
    if (
      !this.editProfile.firstName ||
      !this.editProfile.lastName ||
      !this.editProfile.email
    ) {
      this.showMessage('Veuillez remplir tous les champs obligatoires', 'error');
      return;
    }

    this.isLoading = true;

    this.apiService.updateUser(this.userId, this.editProfile).subscribe({
      next: (res: any) => {
        this.userProfile = { ...res };
        this.isEditMode = false;
        this.isLoading = false;
        this.showMessage('Profil mis à jour avec succès !', 'success');
      },
      error: () => {
        this.isLoading = false;
        this.showMessage('Erreur lors de la mise à jour.', 'error');
      }
    });
  }

  confirmDelete() {
    this.showDeleteModal = true;
  }

  cancelDelete() {
    this.showDeleteModal = false;
  }

  deleteProfile() {
    this.isLoading = true;

    this.apiService.deleteUser(this.userId).subscribe({
      next: () => {
        this.isLoading = false;
        this.showDeleteModal = false;
        localStorage.removeItem('access_token');
        this.router.navigate(['/login']);
      },
      error: () => {
        this.isLoading = false;
        this.showDeleteModal = false;
        this.showMessage('Impossible de supprimer le profil.', 'error');
      }
    });
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
