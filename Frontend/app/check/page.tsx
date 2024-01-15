import { HeroSection } from '@/ui/HeroSection-Explore'
import { RadiusButton } from '@/ui/RadiusButton'
import { Checkbox, Text, Wrap } from "@yamada-ui/react"
import Image from 'next/image'

const inshoData = [
  '好き・嫌い',
  '美しい・汚い',
  'かっこいい・かわいい',
  '明るい・暗い',
  '大人っぽい・子供っぽい',
  '派手・地味',
] 

export default function Check(){
  return(
  <div>
    <HeroSection/>
    <div className='flex flex-col gap-12 text-xl pl-20 pt-20'>
    <Text>以下から検索したい印象を選択</Text>
      {inshoData.map((item) => (
        // eslint-disable-next-line react/jsx-key
        <Checkbox size="lg">{item}</Checkbox>
      ))}
    </div>
    <div className='flex justify-center pt-12'>
      <RadiusButton href='/slide' className='bg-gradient-to-tr from-blue-400 to-pink-300'>
        <Image src="/next_grad.svg" alt='ゲーミングNEXT' width={100} height={100}/>
      </RadiusButton>
    </div>
  </div>
  )
}