import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ApiService } from '../../../services/api.service';
import { ActivatedRoute } from '@angular/router';
import jsPDF from 'jspdf';
import autoTable from 'jspdf-autotable';

interface Presence {
  Type: string;
  Timestamp: string;
}

@Component({
  selector: 'app-employee-presence',
  standalone: true,
  imports: [CommonModule],
  template: `
  <section class="stats-container">

    <h1>Présences de {{ employeeName }}</h1>

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

        <div class="kpi-card neutral">
          <div class="kpi-title">🕒 Aujourd'hui</div>
          <div class="kpi-value">{{ formatHours(todayWorkedHours) }}</div>
          <div class="kpi-sub">heures travaillées</div>
        </div>

        <div class="kpi-card neutral">
          <div class="kpi-title">⏳ Restant (8h)</div>
          <div class="kpi-value">{{ formatHours(todayRemainingHours) }}</div>
        </div>

        <div class="kpi-card" [ngClass]="weeklyStatus">
          <div class="kpi-title">📆 Cette semaine</div>
          <div class="kpi-value">{{ formatHours(weeklyWorkHours) }}</div>
          <div class="kpi-sub">/ 35 h</div>
        </div>

        <div class="kpi-card neutral">
          <div class="kpi-title">🗓 Ce mois-ci</div>
          <div class="kpi-value">{{ formatHours(monthlyWorkHours) }}</div>
          <div class="kpi-sub">heures travaillées</div>
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
  styleUrls: ['./manager_check_presence.component.css']
})
export class EmployeePresenceComponent implements OnInit {

  employeeId!: string;
  employeeName = 'Employé';
  presences: Presence[] = [];
  loading = true;

  totalWorkHours = 0;
  totalDays = 0;
  averageHoursPerDay = 0;

  todayWorkedHours = 0;
  todayRemainingHours = 0;

  weeklyWorkHours = 0;
  monthlyWorkHours = 0;

  readonly goalHoursPerDay = 8;
  readonly weeklyGoalHours = 35;

  private presenceByDay: { [key: string]: Presence[] } = {};

  constructor(
    private api: ApiService,
    private route: ActivatedRoute
  ) { }

  ngOnInit(): void {
    this.employeeId = this.route.snapshot.paramMap.get('id') || '';
    if (this.employeeId) {
      this.loadEmployeePresences();
    }
  }

  get productivityLevel(): 'low' | 'medium' | 'high' {
    if (this.averageHoursPerDay < 4) return 'low';
    if (this.averageHoursPerDay < 7) return 'medium';
    return 'high';
  }

  get weeklyStatus(): 'low' | 'medium' | 'high' {
    if (this.weeklyWorkHours < 25) return 'low';
    if (this.weeklyWorkHours < this.weeklyGoalHours) return 'medium';
    return 'high';
  }

  loadEmployeePresences(): void {
    this.loading = true;

    this.api.getUserPresencesForManager(this.employeeId).subscribe({
      next: (res: any) => {
        this.employeeName = res.userName || 'Employé';
        this.presences = Array.isArray(res.presences) ? res.presences : [];
        this.calculateStats();
        this.loading = false;
      },
      error: () => {
        this.presences = [];
        this.loading = false;
      }
    });
  }

  private calculateStats(): void {

    if (!this.presences.length) return;

    const MS_PER_HOUR = 3600000;
    const presenceByWeek: { [key: string]: Presence[] } = {};
    const presenceByMonth: { [key: string]: Presence[] } = {};

    this.presenceByDay = {};

    this.presences.forEach(p => {
      const date = new Date(p.Timestamp);
      const dayKey = date.toISOString().slice(0, 10);
      const weekKey = this.getISOWeek(date);
      const monthKey = this.getMonthKey(date);

      if (!this.presenceByDay[dayKey]) this.presenceByDay[dayKey] = [];
      if (!presenceByWeek[weekKey]) presenceByWeek[weekKey] = [];
      if (!presenceByMonth[monthKey]) presenceByMonth[monthKey] = [];

      this.presenceByDay[dayKey].push(p);
      presenceByWeek[weekKey].push(p);
      presenceByMonth[monthKey].push(p);
    });

    let totalHours = 0;
    let daysCount = 0;

    for (const day in this.presenceByDay) {
      const dayHours = this.calculateHours(this.presenceByDay[day], MS_PER_HOUR);
      totalHours += dayHours;
      daysCount++;

      if (day === new Date().toISOString().slice(0, 10)) {
        this.todayWorkedHours = dayHours;
        this.todayRemainingHours = Math.max(this.goalHoursPerDay - dayHours, 0);
      }
    }

    const currentWeek = this.getISOWeek(new Date());
    const currentMonth = this.getMonthKey(new Date());

    this.weeklyWorkHours = presenceByWeek[currentWeek]
      ? this.calculateHours(presenceByWeek[currentWeek], MS_PER_HOUR)
      : 0;

    this.monthlyWorkHours = presenceByMonth[currentMonth]
      ? this.calculateHours(presenceByMonth[currentMonth], MS_PER_HOUR)
      : 0;

    this.totalWorkHours = totalHours;
    this.totalDays = daysCount;
    this.averageHoursPerDay = totalHours / daysCount;
  }

  private calculateHours(presences: Presence[], ms: number): number {
    const sorted = presences.sort(
      (a, b) => new Date(a.Timestamp).getTime() - new Date(b.Timestamp).getTime()
    );

    let hours = 0;
    for (let i = 0; i < sorted.length - 1; i += 2) {
      hours +=
        (new Date(sorted[i + 1].Timestamp).getTime() -
          new Date(sorted[i].Timestamp).getTime()) / ms;
    }
    return hours;
  }

  private getISOWeek(date: Date): string {
    const d = new Date(Date.UTC(date.getFullYear(), date.getMonth(), date.getDate()));
    const dayNum = d.getUTCDay() || 7;
    d.setUTCDate(d.getUTCDate() + 4 - dayNum);
    const yearStart = new Date(Date.UTC(d.getUTCFullYear(), 0, 1));
    const weekNo = Math.ceil((((d.getTime() - yearStart.getTime()) / 86400000) + 1) / 7);
    return `${d.getUTCFullYear()}-W${weekNo}`;
  }

  private getMonthKey(date: Date): string {
    return `${date.getFullYear()}-${(date.getMonth() + 1)
      .toString()
      .padStart(2, '0')}`;
  }

  formatHours(hours: number): string {
    const h = Math.floor(hours);
    const m = Math.round((hours - h) * 60);
    return `${h} h ${m.toString().padStart(2, '0')} min`;
  }

  exportToPDF(): void {
    const doc = new jsPDF();
    doc.text(`Présences de ${this.employeeName}`, 14, 20);

    autoTable(doc, {
      startY: 30,
      head: [['Type', 'Date', 'Heure']],
      body: this.presences.map(p => [
        p.Type,
        new Date(p.Timestamp).toLocaleDateString('fr-FR'),
        new Date(p.Timestamp).toLocaleTimeString('fr-FR')
      ])
    });

    doc.save(`presences-${this.employeeName}.pdf`);
  }

  exportToCSV(): void {
    const rows = this.presences.map(p =>
      `${p.Type};${new Date(p.Timestamp).toLocaleDateString('fr-FR')};${new Date(p.Timestamp).toLocaleTimeString('fr-FR')}`
    );

    const blob = new Blob([rows.join('\n')], { type: 'text/csv;charset=utf-8' });
    const link = document.createElement('a');
    link.href = URL.createObjectURL(blob);
    link.download = `presences-${this.employeeName}.csv`;
    link.click();
  }
}
