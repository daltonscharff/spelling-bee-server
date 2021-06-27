import { IsString, IsInt } from 'class-validator';

export class CreateWordDto {
    @IsString()
    word: string;

    @IsInt()
    pointValue: number;

    @IsString()
    definition: string;

    @IsString()
    partOfSpeech: string;

    @IsString()
    synonym: string;
}
