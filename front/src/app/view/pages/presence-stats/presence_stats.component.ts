import {
  Component,
  OnInit
} from '@angular/core';
import { CommonModule } from '@angular/common';
import { ApiService } from '../../../services/api.service';
import jsPDF from 'jspdf';
import autoTable from 'jspdf-autotable';

interface Presence {
  Type: string;
  Timestamp: string;
}

@Component({
  selector: 'app-presence-stats',
  standalone: true,
  imports: [CommonModule],
  template: `
  <section class="stats-container">

    <h1>Statistiques de présence</h1>

    <div *ngIf="loading">Chargement des données...</div>

    <div *ngIf="!loading && presences.length === 0">
      Aucune présence enregistrée.
    </div>

    <div *ngIf="!loading && presences.length > 0">

      <!-- KPI -->
      <div class="kpi-grid">

        <div class="kpi-card neutral">
          <div class="kpi-title">⏱ Heures totales</div>
          <div class="kpi-value">{{ formatHours(totalWorkHours) }}</div>
          <div class="kpi-sub">temps travaillé</div>
        </div>

        <div class="kpi-card neutral">
          <div class="kpi-title">📅 Jours travaillés</div>
          <div class="kpi-value">{{ totalDays }}</div>
          <div class="kpi-sub">jours</div>
        </div>

        <div class="kpi-card" [ngClass]="productivityLevel">
          <div class="kpi-title">📊 Moyenne / jour</div>
          <div class="kpi-value">{{ formatHours(averageHoursPerDay) }}</div>
          <div class="kpi-sub">productivité</div>
        </div>

      </div>

      <!-- Export -->
      <div class="export-buttons">
        <button class="export-btn" (click)="exportToPDF()">📄 Exporter en PDF</button>
        <button class="export-btn csv" (click)="exportToCSV()">📊 Exporter en CSV</button>
      </div>

      <!-- Table -->
      <h2>Historique des présences</h2>
      <table>
        <thead>
          <tr>
            <th>Type</th>
            <th>Date</th>
            <th>Heure</th>
          </tr>
        </thead>
        <tbody>
          <tr *ngFor="let p of presences">
            <td>{{ p.Type }}</td>
            <td>{{ p.Timestamp | date:'dd/MM/yyyy' }}</td>
            <td>{{ p.Timestamp | date:'HH:mm:ss' }}</td>
          </tr>
        </tbody>
      </table>

    </div>

  </section>
  `,
  styleUrls: ['./presence_stats.component.css']
})
export class PresenceStatsComponent implements OnInit {

  presences: Presence[] = [];
  loading = true;

  totalWorkHours = 0;
  totalDays = 0;
  averageHoursPerDay = 0;

  private presenceByDay: { [key: string]: Presence[] } = {};

  constructor(private api: ApiService) { }

  ngOnInit() {
    this.api.getPresence().subscribe({
      next: (res: any) => {
        this.presences = res.user || [];
        this.calculateStats();
        this.loading = false;
      },
      error: err => {
        console.error('Erreur récupération présences', err);
        this.loading = false;
      }
    });
  }

  get productivityLevel(): 'low' | 'medium' | 'high' {
    if (this.averageHoursPerDay < 4) return 'low';
    if (this.averageHoursPerDay < 7) return 'medium';
    return 'high';
  }

  private calculateStats() {
    this.presenceByDay = {};
    const MS_PER_HOUR = 3600000;

    this.presences.forEach(p => {
      const day = new Date(p.Timestamp).toISOString().slice(0, 10);
      if (!this.presenceByDay[day]) this.presenceByDay[day] = [];
      this.presenceByDay[day].push(p);
    });

    let totalHours = 0;
    let daysCount = 0;

    for (const day in this.presenceByDay) {
      const sorted = this.presenceByDay[day].sort(
        (a, b) =>
          new Date(a.Timestamp).getTime() -
          new Date(b.Timestamp).getTime()
      );

      for (let i = 0; i < sorted.length - 1; i += 2) {
        totalHours +=
          (new Date(sorted[i + 1].Timestamp).getTime() -
            new Date(sorted[i].Timestamp).getTime()) / MS_PER_HOUR;
      }

      daysCount++;
    }

    this.totalWorkHours = totalHours;
    this.totalDays = daysCount;
    this.averageHoursPerDay =
      daysCount > 0 ? totalHours / daysCount : 0;
  }

  formatHours(hours: number): string {
    const h = Math.floor(hours);
    const m = Math.round((hours - h) * 60);
    return `${h} h ${m.toString().padStart(2, '0')} min`;
  }

  exportToPDF() {
    const doc = new jsPDF();
    doc.setFontSize(16);
    doc.text('Statistiques de présence', 14, 20);

    doc.setFontSize(12);
    doc.text(`Heures totales : ${this.formatHours(this.totalWorkHours)}`, 14, 35);
    doc.text(`Jours travaillés : ${this.totalDays}`, 14, 42);
    doc.text(`Moyenne / jour : ${this.formatHours(this.averageHoursPerDay)}`, 14, 49);

    autoTable(doc, {
      startY: 60,
      head: [['Type', 'Date', 'Heure']],
      body: this.presences.map(p => [
        p.Type,
        new Date(p.Timestamp).toLocaleDateString('fr-FR'),
        new Date(p.Timestamp).toLocaleTimeString('fr-FR')
      ])
    });

    doc.save('statistiques-presence.pdf');
  }

  exportToCSV() {
    const rows = this.presences.map(p =>
      `${p.Type};${new Date(p.Timestamp).toLocaleDateString('fr-FR')};${new Date(
        p.Timestamp
      ).toLocaleTimeString('fr-FR')}`
    );

    const blob = new Blob([rows.join('\n')], {
      type: 'text/csv;charset=utf-8'
    });

    const link = document.createElement('a');
    link.href = URL.createObjectURL(blob);
    link.download = 'statistiques-presence.csv';
    link.click();
  }
}
