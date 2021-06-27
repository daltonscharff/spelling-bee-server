import { Entity, Column, PrimaryGeneratedColumn } from "typeorm";

@Entity('words')
export class Word {
    @PrimaryGeneratedColumn('uuid')
    id: string;

    @Column({ length: 64 })
    word: string;

    @Column({ type: 'smallint', nullable: true })
    pointValue: number;

    @Column({ type: 'text', nullable: true })
    definition: string;

    @Column({ length: 32, nullable: true })
    partOfSpeech: string;

    @Column({ length: 64, nullable: true})
    synonym: string;
}