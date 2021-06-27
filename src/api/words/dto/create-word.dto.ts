import { IsString, IsInt, IsPositive, IsOptional, IsAlpha } from 'class-validator';

export class CreateWordDto {
    @IsString()
    @IsAlpha()
    word: string;

    @IsInt()
    @IsPositive()
    @IsOptional()
    pointValue: number;

    @IsString()
    @IsOptional()
    definition: string;

    @IsString()
    @IsOptional()
    partOfSpeech: string;

    @IsString()
    @IsOptional()
    synonym: string;
}
