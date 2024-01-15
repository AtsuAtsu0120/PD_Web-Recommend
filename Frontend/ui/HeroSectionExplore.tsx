//ヒーローセクション.explore
import { Center,Box,VStack } from "@yamada-ui/react";

export function HeroSectionExplore() {
    return(
        <div>
            
            <Box position="relative"boxSize="2xs"text="40"fontWeight="550"fontStyle="oblique"bg="primary"p="md"color="white"_after={{content: "'ASA'",position: "absolute", top: "5", left: "15",bg: "secondary",color: "#EBEBEB",}}>

            </Box>

            <Center h="300" rounded="md" bg="#064E6F">
            <VStack>

                <Center text="32" fontWeight="550" rounded="md" color="#EBEBEB">
                 Explore
                </Center>
                <Center text="15" rounded="md" color="#ECECEC">
                 サイトを見つける
                </Center>

            </VStack>
            </Center>
              
        </div>
    )};