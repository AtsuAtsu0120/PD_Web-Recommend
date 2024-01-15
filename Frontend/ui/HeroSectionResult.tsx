//ヒーローセクション.result
import {Center,Box,VStack } from "@yamada-ui/react";

export function HeroSectionResult() {
    return(
        <div>
            
            <Box position="relative"boxSize="2xs"text="40"fontWeight="550"fontStyle="oblique"bg="primary"p="md"color="#141414"_after={{content: "'ASA'",position: "absolute", top: "5", left: "15",bg: "secondary",color: "#474747",}}>

            </Box>

            <Center h="300" rounded="md" bg="#FFFFFF">
            <VStack>

                <Center text="32" fontWeight="800" rounded="md" color="#474747">
                 Result
                </Center>
                <Center text="15" rounded="md" color="#212121">
                 見つかったサイト
                </Center>

            </VStack>
            </Center>
              
        </div>
    );
  }