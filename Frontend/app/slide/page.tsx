import { HeroSection } from '@/ui/HeroSection-Explore'
import { RadiusButton } from '@/ui/RadiusButton'
import { Checkbox, Text, Wrap, Slider } from "@yamada-ui/react"
import Image from 'next/image'

type SliderText = {
  positive: string;
  negative: string;
}

const inshoData: SliderText[] = [
  {
    positive: '好き',
    negative: '嫌い',
  },
  {
    positive: '美しい',
    negative: '汚い',
  },
  {
    positive: 'かっこいい',
    negative: 'かわいい',
  },
  {
    positive: '明るい',
    negative: '暗い',
  },
  {
    positive: '大人っぽい',
    negative: '子供っぽい',
  },
  {
    positive: '派手',
    negative: '地味',
  },
] 

export default function Slidef(){
  return(
  <div>
    <HeroSection/>
    <div className='flex flex-col gap-12 text-xl pt-20 justify-center items-center'>
    <Text className="text-center">それぞれの印象の程度を入力</Text>
      {inshoData.map((item) => (
        // eslint-disable-next-line react/jsx-key
        <div className='flex justify-center pb-5'>
          <div className='w-96 flex justify-between'>
            <div>{item.positive}</div>
            <div>{item.negative}</div>
          </div>
          <Slider/>
        </div>
      ))}
    </div>
    <div className='flex justify-center py-20 gap-12'>
      <RadiusButton href='/check' className='text-3xl py-1'>
        Back
      </RadiusButton>
      <RadiusButton href='/result' className='bg-gradient-to-tr from-blue-400 to-pink-300'>
        <Image src="/go_grad.svg" alt='ゲーミングGO' width={100} height={100}/>
      </RadiusButton>
    </div>
  </div>
  )
}