import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Profile } from './profile.model';

@Injectable({ providedIn: 'root' })
export class APIService {
  profile: any;

  constructor(private httpClient: HttpClient) {}

  getProfile() {
    this.httpClient
      .get('http://144.126.154.126:10000/getbitches')
      .subscribe((data) => {
        console.log(data);
        this.profile = {
          
        };
      });
  }
}
