"use client"

import { Box_a } from '@/ui/Box'
import { HeroSection } from '@/ui/HeroSection-Explore'
import { RadiusButton } from '@/ui/RadiusButton'
import { Checkbox, Text, Wrap, Slider, SliderTrack, SliderThumb, Button, Modal, ModalHeader, ModalBody, ModalFooter, ModalOverlay } from "@yamada-ui/react"
import Image from 'next/image'
import { FormEvent, useState } from 'react'

type SliderText = {
  name: string
  positive: string;
  negative: string;
}

const inshoData: SliderText[] = [
  {
    name: "like",
    positive: '好き',
    negative: '嫌い',
  },
  {
    name: "beautiful",
    positive: '美しい',
    negative: '汚い',
  },
  {
    name: "smart",
    positive: 'かっこいい',
    negative: 'かわいい',
  },
  {
    name: "bright",
    positive: '明るい',
    negative: '暗い',
  },
  {
    name: 'adult',
    positive: '大人っぽい',
    negative: '子供っぽい',
  },
  {
    name: "flashy",
    positive: '派手',
    negative: '地味',
  },
] 

export default function Slidef() {
  const [isOpen, setIsOpen] = useState(false)
  const onClose = () => {
    setIsOpen(false)
  }
  type BoxDeta = {
    imgsrc: string;
    href: string;
    title: string;
  }
  let [boxData, setBoxData] = useState([
    {
      imgsrc: '/cirkitHP.png',
      href: 'https://cirkit.jp',
      title: '株式会社CirKit',
    },
  ])

  async function onSubmit(e: FormEvent<HTMLFormElement>) {
    e.preventDefault()

    const formData = new FormData(e.currentTarget);

    formData.forEach(element => {
      console.log(element)
    });
    const response = await fetch("http://localhost:1323/ranking", {
      method: 'POST',
      body: formData,
      mode: "cors",
      credentials: "same-origin"
    })

    const data = await response.json();
    boxData = new Array<BoxDeta>();

    data.forEach((element: { url: string; name: string, filePath: string }) => {
      let fileName = element.filePath;
      fileName = fileName.replaceAll(" ", "_");
      fileName = fileName.replaceAll("|", "");
      fileName = fileName.replaceAll(":", "");
      fileName = fileName.replaceAll("?", "");

      let filePath = '/images' + fileName;

      console.log(filePath);

      boxData.push({
        imgsrc: filePath,
        href: element.url,
        title: element.name
      })
    });
    console.log(data)
    setBoxData(boxData)
    setIsOpen(true)
  }
  return(
    <div>
      <HeroSection/>
      <form onSubmit={onSubmit}>
        <div className='flex flex-col gap-12 text-xl pt-20 justify-center items-center'>
        <Text className="text-center">それぞれの印象の程度を入力</Text>
          {inshoData.map((item) => (
            // eslint-disable-next-line react/jsx-key
            <div className='flex-col justify-center pb-5' key={item.name}>
              <div className='w-96 flex justify-between'>
                <div>{item.positive}</div>
                <div>{item.negative}</div>
              </div>
              <input name={item.name} type="range" min="-3" max="3" className="w-full h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer dark:bg-gray-700"/>
            </div>
          ))}
        </div>
        <div className='flex justify-center py-20 gap-12'>
          <RadiusButton href='/check' className='text-3xl py-1'>
            Back
          </RadiusButton>
          <Button className="p-1 rounded-2xl font-bold" type='submit'>
            <div className="bg-[#042f43] p-7 rounded-2xl">
              <Image src="/go_grad.svg" alt='ゲーミングGO' width={100} height={100}/>
            </div>
          </Button>
        </div>
        </form>
        <Modal isOpen={isOpen} onClose={onClose}>
          <ModalBody>
          <div className='bg-[#042f43]'>
            <div className='flex flex-col gap-12 text-xl pt-20 justify-center items-center'>
            <Text className="text-center text-slate-100 font-bold text-4xl">結果</Text>
              {boxData.map((item) => (
                // eslint-disable-next-line react/jsx-key
                <Box_a imgsrc={item.imgsrc} href={item.href} title={item.title}/>
              ))}
            </div>
          </div>
          </ModalBody>

          <ModalFooter>
            <Button variant="ghost" onClick={onClose}>
              とじる
            </Button>
            <Button colorScheme="primary">Wikipadia</Button>
          </ModalFooter>
        </Modal>
    </div>
  )
}