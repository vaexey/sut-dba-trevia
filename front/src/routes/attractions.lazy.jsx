import { createLazyFileRoute, useRouterState } from "@tanstack/react-router";

export const Route = createLazyFileRoute("/attractions")({
  component: AttractionsPage,
});

function AttractionsPage() {
  const search = useRouterState({ select: (s) => s.location.search });
  const { id, category } = search;

  console.log(id, category);
}
