import { Component, OnInit } from '@angular/core';
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
      <p>Total d'heures travaillées : {{ formatHours(totalWorkHours) }}</p>
      <p>Nombre de jours : {{ totalDays }}</p>
      <p>Moyenne par jour : {{ formatHours(averageHoursPerDay) }}</p>

      <div class="export-buttons">
        <button class="export-btn" (click)="exportToPDF()">📄 Exporter en PDF</button>
        <button class="export-btn csv" (click)="exportToCSV()">📊 Exporter en CSV</button>
      </div>

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

  constructor(private api: ApiService) { }

  ngOnInit() {
    this.api.getPresence().subscribe({
      next: (res: any) => {
        this.presences = res.user || [];
        this.calculateStats();
        this.loading = false;
      },
      error: (err) => {
        console.error('Erreur récupération présences', err);
        this.loading = false;
      }
    });
  }

  private calculateStats() {
    if (this.presences.length === 0) return;

    const presenceByDay: { [key: string]: Presence[] } = {};
    const MS_PER_HOUR = 3600000;
    this.presences.forEach(p => {
      const day = new Date(p.Timestamp).toISOString().slice(0, 10);
      if (!presenceByDay[day]) presenceByDay[day] = [];
      presenceByDay[day].push(p);
    });

    let totalHours = 0;
    let daysCount = 0;

    for (const day in presenceByDay) {
      const dayPresences = presenceByDay[day].sort((a, b) =>
        new Date(a.Timestamp).getTime() - new Date(b.Timestamp).getTime()
      );

      for (let i = 0; i < dayPresences.length - 1; i += 2) {
        const start = new Date(dayPresences[i].Timestamp).getTime();
        const end = new Date(dayPresences[i + 1].Timestamp).getTime();
        totalHours += (end - start) / MS_PER_HOUR;
      }
      daysCount++;
    }

    this.totalWorkHours = totalHours;
    this.totalDays = daysCount;
    this.averageHoursPerDay = daysCount > 0 ? totalHours / daysCount : 0;
  }

  formatHours(hours: number): string {
    const h = Math.floor(hours);
    const m = Math.round((hours - h) * 60);
    return `${h} h ${m.toString().padStart(2, '0')} min`;
  }

  //Export PDF
  exportToPDF() {
    const doc = new jsPDF();

    doc.setFontSize(16);
    doc.text('Statistiques de présence', 14, 20);
    doc.setFontSize(12);

    doc.text(`Total d'heures travaillées : ${this.formatHours(this.totalWorkHours)}`, 14, 35);
    doc.text(`Nombre de jours : ${this.totalDays}`, 14, 42);
    doc.text(`Moyenne par jour : ${this.formatHours(this.averageHoursPerDay)}`, 14, 49);

    const tableData = this.presences.map(p => [
      p.Type,
      new Date(p.Timestamp).toLocaleDateString('fr-FR'),
      new Date(p.Timestamp).toLocaleTimeString('fr-FR')
    ]);

    autoTable(doc, {
      startY: 60,
      head: [['Type', 'Date', 'Heure']],
      body: tableData,
    });

    const date = new Date().toISOString().slice(0, 10);
    doc.save(`statistiques-presence-${date}.pdf`);
  }


  exportToCSV() {
    const date = new Date().toISOString().slice(0, 10);

    const summary = [
      `Statistiques de présence (${date})`,
      `Total d'heures travaillées;${this.formatHours(this.totalWorkHours)}`,
      `Nombre de jours;${this.totalDays}`,
      `Moyenne par jour;${this.formatHours(this.averageHoursPerDay)}`,
      '', // ligne vide
      'Historique des présences',
      'Type;Date;Heure'
    ];

    const rows = this.presences.map(p => {
      const d = new Date(p.Timestamp);
      const dateStr = d.toLocaleDateString('fr-FR');
      const timeStr = d.toLocaleTimeString('fr-FR');
      return `${p.Type};${dateStr};${timeStr}`;
    });

    const csvContent = [...summary, ...rows].join('\n');

    const blob = new Blob([csvContent], { type: 'text/csv;charset=utf-8;' });
    const url = URL.createObjectURL(blob);
    const link = document.createElement('a');

    link.href = url;
    link.download = `statistiques-presence-${date}.csv`;
    link.click();

    URL.revokeObjectURL(url);
  }
}
