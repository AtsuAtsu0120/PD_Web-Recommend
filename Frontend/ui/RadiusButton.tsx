import { Button } from "@yamada-ui/react"
import { Link } from "@yamada-ui/react"
import { twMerge } from 'tailwind-merge'

type RadiusButtonProps = {
    children: React.ReactNode;
    href: string;
    className?: string;
}

export function RadiusButton({children, href, className}: RadiusButtonProps){
    const BorderClassName = twMerge(`p-1 rounded-2xl font-bold ${children == undefined ? '' : 'bg-slate-50'}`,className);
    return (
        <Link href={href}>
        <Button className={BorderClassName}>
            <div className="bg-[#042f43] p-7 rounded-2xl">
            {children}
            </div>
        </Button>
        </Link>
    )
}