import SystemService from '../service/SystemService';

class SystemController {
  private service: SystemService = new SystemService();
}

export const systemController = new SystemController();

export const Routes = [];
