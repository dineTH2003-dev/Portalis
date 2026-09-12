export interface ServerConfig {
  port: number;
  jwtAccessSecret: string;
  jwtRefreshSecret: string;
  tunnelGrantSecret: string;
  internalGatewaySecret: string;
  gatewayWSURL: string;
}

export function loadEnv(): ServerConfig {
  return {
    port: Number(process.env.PORT) || 4310,
    jwtAccessSecret: process.env.JWT_ACCESS_SECRET || "change_me_jwt_access_secret_32_chars",
    jwtRefreshSecret: process.env.JWT_REFRESH_SECRET || "change_me_jwt_refresh_secret_32_chars",
    tunnelGrantSecret: process.env.TUNNEL_GRANT_SECRET || "dev_secret_grant_change_in_prod",
    internalGatewaySecret: process.env.INTERNAL_GATEWAY_SECRET || "dev_secret_gateway_change_in_prod",
    gatewayWSURL: process.env.GATEWAY_WS_URL || "ws://localhost:9000/v1/tunnel/ws",
  };
}
