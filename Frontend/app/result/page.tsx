import { Box_a } from '@/ui/Box'
import { HeroSectinon } from '@/ui/HeroSection-Result'
import { RadiusButton } from '@/ui/RadiusButton'
import { Checkbox, Text, Wrap, Slider } from "@yamada-ui/react"
import Image from 'next/image'

type BoxDeta = {
  imgsrc: string;
  href: string;
  title: string;
}

const BoxData: BoxDeta[] = [
  {
    imgsrc: '/cirkitHP.png',
    href: 'https://cirkit.jp',
    title: '株式会社CirKit',
  },
] 

export default function Slidef(){
  return(
  <div className='bg-[#042f43]'>
    <HeroSectinon/>
    <div className='flex flex-col gap-12 text-xl pt-20 justify-center items-center'>
    <Text className="text-center text-slate-100 font-bold text-4xl">結果</Text>
      {BoxData.map((item) => (
        // eslint-disable-next-line react/jsx-key
        <Box_a imgsrc={item.imgsrc} href={item.href} title={item.title}/>
      ))}
    </div>
  </div>
  )
}