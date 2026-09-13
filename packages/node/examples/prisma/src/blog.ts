import { prisma } from "./db";

export async function signUp(email: string) {
  return prisma.user.create({ data: { email } });
}

export async function publish(email: string, title: string) {
  return prisma.$transaction(async (tx) => {
    const author = await tx.user.findUniqueOrThrow({ where: { email } });
    return tx.post.create({ data: { title, authorId: author.id } });
  });
}

export async function postsBy(email: string) {
  return prisma.post.findMany({ where: { author: { email } }, orderBy: { id: "asc" } });
}
