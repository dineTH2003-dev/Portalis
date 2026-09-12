/**
 * Subdomain Service
 * Manages custom vanity subdomains, collision detection, and blocklists.
 */

export interface ReservedSubdomain {
  id: string;
  subdomain: string;
  userId: string;
  createdAt: string;
}

/**
 * Reserve a vanity subdomain.
 * TODO(contributor): [Issue #7]
 *   1. Validate syntax: 3-63 chars, lowercase alphanumeric, hyphens allowed (cannot start or end with hyphen).
 *   2. Check blocked_subdomains table. Return 422 if blocked.
 *   3. Check reserved_subdomains table. Return 409 if already reserved.
 *   4. Check user's max_subdomains quota limit.
 *   5. Insert into reserved_subdomains table.
 */
export async function reserveSubdomain(userId: string, subdomain: string): Promise<ReservedSubdomain> {
  console.log(`[SUBDOMAIN STUB] Reserving ${subdomain} for user ${userId}...`);
  return {
    id: "subdomain-stub-uuid",
    subdomain,
    userId,
    createdAt: new Date().toISOString(),
  };
}

/**
 * List reserved subdomains for a user.
 */
export async function listUserSubdomains(userId: string): Promise<ReservedSubdomain[]> {
  return [
    {
      id: "subdomain-stub-uuid",
      subdomain: "my-app",
      userId,
      createdAt: new Date().toISOString(),
    },
  ];
}
