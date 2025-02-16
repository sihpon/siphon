import Link from "next/link";

interface DetailButtonProps {
  href: string;
}

export default function DetailButton({ href }: DetailButtonProps) {
  return (
    <Link className="btn btn-primary" href={href}>
      詳細
    </Link>
  );
}
