import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({ providedIn: 'root' })
export class ApiService {
    private baseUrl = '/api';

    constructor(private http: HttpClient) { }

    // Authentification
    login(credentials: { Email: string; Password: string }) {
        return this.http.post<{ token: string }>(`${this.baseUrl}/authentification`, credentials);
    }

    // Utilisateur
    getUser(_id: string) {
        return this.http.get(`${this.baseUrl}/user/${_id}`);
    }

    updateUser(_id: string, data: any) {
        return this.http.put(`${this.baseUrl}/user/${_id}`, data);
    }

    deleteUser(_id: string) {
        return this.http.delete(`${this.baseUrl}/user/${_id}`);
    }

    // Équipes
    getAllTeams() {
        return this.http.get<any[]>(`${this.baseUrl}/team`);
    }

    createTeam(data: any) {
        return this.http.post(`${this.baseUrl}/team`, data);
    }

    updateTeam(id: string, data: any) {
        return this.http.put(`${this.baseUrl}/team/${id}`, data);
    }

    deleteTeam(id: string) {
        return this.http.delete(`${this.baseUrl}/team/${id}`);
    }

    // Présence
    postPresence(data: { presence: { Type: string; Timestamp: string } }) {
        return this.http.post(`${this.baseUrl}/presence/create`, data);
    }
}
