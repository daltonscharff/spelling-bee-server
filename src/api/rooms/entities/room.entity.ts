import { Column, Entity, PrimaryGeneratedColumn, Unique } from "typeorm";

@Entity('rooms')
export class Room {
    @PrimaryGeneratedColumn('uuid')
    id: string;

    @Column({ type: 'int' })
    score: number;

    @Column({ length: 128, nullable: true })
    name: string;
}
