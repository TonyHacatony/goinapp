import http from 'k6/http';
import { check, sleep } from 'k6';

// Simple test configuration: 10 users, each making 20 calls
export let options = {
  vus: 10,
  iterations: 200, 
  thresholds: {
    http_req_duration: ['p(95)<1000'],
    http_req_failed: ['rate<0.1'],
  },
};

const BASE_URL = 'http://app:8088';

export default function() {
  let response = http.get(`${BASE_URL}/list`);
  
  check(response, {
    'GET /list status is 200': (r) => r.status === 200,
    'GET /list response time < 1s': (r) => r.timings.duration < 1000,
    'GET /list returns JSON array': (r) => {
      try {
        const data = JSON.parse(r.body);
        return Array.isArray(data);
      } catch (e) {
        return false;
      }
    },
  });

  sleep(0.1);
}
