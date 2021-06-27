import { IsString, IsInt, IsPositive, IsOptional } from 'class-validator';

export class CreateWordDto {
    @IsString()
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
