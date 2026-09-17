package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pgp_cfb_free(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+24))
	m.T0[v4].(func(*base.Module, int32))(m, v3)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		base.MemoryFill(m, l0, int32(0), int32(116))
		F_pfree(m, l0)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			return
		}
	}
}
func F_pgp_create_pkt_reader(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	v7 = F_palloc(m, int32(8))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = l3
		v14 = F_pullf_create(m, l0, int32(_a_F_pgp_create_pkt_reader_0), v7, l1)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			if v14 < int32(0) {
				F_pfree(m, v7)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					return v14
				}
			} else {
				return v14
			}
		}
	}
}
func F_pgp_get_unicode_mode(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	return v2
}
func F_pgp_key_free(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v65 int32
	_ = v65
	if l0 != 0 {
		v10 = int32(24)
		v11 = int32(12)
		v12 = int32(8)
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
		switch v13 - int32(1) {
		case 0, 1, 2:
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v17 = F_pgp_mpi_free(m, v16)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				v24 = int32(24)
				v25 = int32(12)
				v26 = int32(28)
				v27 = int32(32)
				v28 = int32(36)
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0+v25)))
				v31 = F_pgp_mpi_free(m, v30)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					v33 = v26
					v34 = v24
					v37 = v28
					v38 = v27
					v40 = *(*int32)(unsafe.Add(mBase, uint32(l0+v34)))
					v41 = F_pgp_mpi_free(m, v40)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(l0+v33)))
						v45 = F_pgp_mpi_free(m, v44)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return
						} else {
							v48 = *(*int32)(unsafe.Add(mBase, uint32(l0+v38)))
							v49 = F_pgp_mpi_free(m, v48)
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return
							} else {
								v52 = *(*int32)(unsafe.Add(mBase, uint32(l0+v37)))
								v53 = F_pgp_mpi_free(m, v52)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return
								} else {
									base.MemoryFill(m, l0, int32(0), int32(52))
									F_pfree(m, l0)
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return
									} else {
										return
									}
								}
							}
						}
					}
				}
			}
		default:
			base.MemoryFill(m, l0, int32(0), int32(52))
			F_pfree(m, l0)
			mBase = m.M
			v65 = m.ExcPending
			if v65 != 0 {
				return
			} else {
				return
			}
		case 15:
			v33 = v11
			v34 = v12
			v37 = v10
			v38 = v13
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l0+v34)))
			v41 = F_pgp_mpi_free(m, v40)
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return
			} else {
				v44 = *(*int32)(unsafe.Add(mBase, uint32(l0+v33)))
				v45 = F_pgp_mpi_free(m, v44)
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return
				} else {
					v48 = *(*int32)(unsafe.Add(mBase, uint32(l0+v38)))
					v49 = F_pgp_mpi_free(m, v48)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return
					} else {
						v52 = *(*int32)(unsafe.Add(mBase, uint32(l0+v37)))
						v53 = F_pgp_mpi_free(m, v52)
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return
						} else {
							base.MemoryFill(m, l0, int32(0), int32(52))
							F_pfree(m, l0)
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			}
		case 16:
			v24 = v11
			v25 = v12
			v26 = int32(16)
			v27 = int32(20)
			v28 = v10
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l0+v25)))
			v31 = F_pgp_mpi_free(m, v30)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return
			} else {
				v33 = v26
				v34 = v24
				v37 = v28
				v38 = v27
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l0+v34)))
				v41 = F_pgp_mpi_free(m, v40)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(l0+v33)))
					v45 = F_pgp_mpi_free(m, v44)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return
					} else {
						v48 = *(*int32)(unsafe.Add(mBase, uint32(l0+v38)))
						v49 = F_pgp_mpi_free(m, v48)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return
						} else {
							v52 = *(*int32)(unsafe.Add(mBase, uint32(l0+v37)))
							v53 = F_pgp_mpi_free(m, v52)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return
							} else {
								base.MemoryFill(m, l0, int32(0), int32(52))
								F_pfree(m, l0)
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				}
			}
		}
	} else {
		return
	}
}
func F_pgp_mpi_write(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v10 = int32(8)
	v14 = v9<<(uint(v10)%32) | int32(base.Ui32(v9)>>(uint(v10)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+14)) = uint16(v14)
	v19 = F_pushf_write(m, l0, v7+int32(14), int32(2))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		if int32(0) <= v19 {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v27 = F_pushf_write(m, l0, v25, v26)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				v29 = v27
				m.G0 = v7 + int32(16)
				return v29
			}
		} else {
			v29 = v19
			m.G0 = v7 + int32(16)
			return v29
		}
	}
}
func F_pgp_rsa_encrypt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v12 = int32(-109)
	v13 = F_mpi_check(m, l1)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		if v13 == int32(0) {
			v58 = v12
			return v58
		} else {
			v19 = F_mpi_check(m, v9)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				if v19 == int32(0) {
					v58 = v12
					return v58
				} else {
					v23 = F_mpi_check(m, v10)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						if v23 == int32(0) {
							v58 = v12
							return v58
						} else {
							v27 = int32(1)
							if v11 <= v27 {
								v30 = v27
							} else {
								v30 = v11
							}
							v31 = F_palloc(m, v30)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int32(0)
							} else {
								v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								v36 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
								v37 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
								v38 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
								v39 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
								v40 = m.Env.Pgmem_bn_op(m, int32(1), v34, v35, v36, v37, v38, v39, v31, v11)
								mBase = m.M
								if v40 < int32(0) {
									v51 = int32(-109)
									if v31 == int32(0) {
										v58 = v51
										return v58
									} else {
										if v30 != 0 {
											base.MemoryFill(m, v31, int32(0), v30)
										} else {
										}
										F_pfree(m, v31)
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return int32(0)
										} else {
											v58 = v51
											return v58
										}
									}
								} else {
									v44 = F_bytes_to_mpi(m, v31, v40)
									mBase = m.M
									v45 = m.ExcPending
									if v45 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l2))) = v44
										if v44 != 0 {
											v49 = int32(0)
										} else {
											v49 = int32(-109)
										}
										v51 = v49
										if v31 == int32(0) {
											v58 = v51
											return v58
										} else {
											if v30 != 0 {
												base.MemoryFill(m, v31, int32(0), v30)
											} else {
											}
											F_pfree(m, v31)
											mBase = m.M
											v57 = m.ExcPending
											if v57 != 0 {
												return int32(0)
											} else {
												v58 = v51
												return v58
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_pgp_s2k_process(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v349 int32
	_ = v349
	var v371 int32
	_ = v371
	var v393 int32
	_ = v393
	var v400 int32
	_ = v400
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	v5 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(80)
	m.G0 = v19
	v23 = l1 - int32(2)
	if base.B2i32(base.Ui32(int32(8)) < base.Ui32(v23))|base.B2i32(int32(base.Ui32(int32(487))>>(uint(v23)%32))&int32(1) == v5) != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+43)) = uint8(v40)
	if v40&int32(255) == int32(0) {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	v40 = int32(0)
	goto L4
L3:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v23<<(uint(int32(2))%32))+uint32(_c_F_pgp_s2k_process[0])))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	v40 = v39
	goto L4
L4:
	;
	goto L1
L5:
	;
	m.G0 = v19 + int32(80)
	return v420
L6:
	;
	v420 = int32(-103)
	goto L5
L7:
	;
	goto L8
L8:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	v50 = F_pgp_load_digest(m, v47, v19+int32(12))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	if v50 < int32(0) {
		v420 = v50
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	switch v57 {
	case 0:
		goto L17
	case 1:
		goto L16
	default:
		v400 = int32(-121)
		goto L12
	case 3:
		goto L15
	}
L12:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v415)+20))
	m.T0[v416].(func(*base.Module, int32))(m, v415)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L9
	} else {
		goto L115
	}
L13:
	;
	v393 = int32(0)
	goto L112
L14:
	;
	v371 = int32(0)
	goto L108
L15:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
	v181 = m.T0[v180].(func(*base.Module, int32) int32)(m, v179)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L9
	} else {
		goto L61
	}
L16:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	v117 = m.T0[v116].(func(*base.Module, int32) int32)(m, v115)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L9
	} else {
		goto L39
	}
L17:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v60 = m.T0[v59].(func(*base.Module, int32) int32)(m, v58)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+43)))
	if v62 == int32(0) {
		goto L13
	} else {
		goto L19
	}
L19:
	;
	v68 = int32(0)
	v69 = v62
	v76 = l0 + int32(11)
	goto L20
L20:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
	m.T0[v84].(func(*base.Module, int32))(m, v58)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L9
	} else {
		goto L22
	}
L22:
	;
	if v68 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	if v68 != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
	m.T0[v96].(func(*base.Module, int32, int32, int32))(m, v58, l2, l3)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L9
	} else {
		goto L30
	}
L26:
	;
	base.MemoryFill(m, v19+int32(16), int32(0), v68)
	goto L28
L27:
	;
	goto L28
L28:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
	m.T0[v93].(func(*base.Module, int32, int32, int32))(m, v58, v19+int32(16), v68)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L9
	} else {
		goto L29
	}
L29:
	;
	goto L25
L30:
	;
	v100 = v19 + int32(16)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v58)+16))
	m.T0[v101].(func(*base.Module, int32, int32))(m, v58, v100)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L9
	} else {
		goto L31
	}
L31:
	;
	if base.Ui32(v69) <= base.Ui32(v60) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	if v69 == int32(0) {
		goto L13
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	if v60 != 0 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	base.MemoryCopy(m, v76, v100, v69)
	goto L13
L36:
	;
	base.MemoryCopy(m, v76, v19+int32(16), v60)
	goto L38
L37:
	;
	goto L38
L38:
	;
	v68 = v68 + int32(1)
	v69 = v69 - v60
	v76 = v60 + v76
	goto L20
L39:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+43)))
	if v119 == int32(0) {
		goto L14
	} else {
		goto L40
	}
L40:
	;
	v127 = int32(0)
	v128 = v119
	v135 = l0 + int32(11)
	goto L41
L41:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v115)+8))
	m.T0[v143].(func(*base.Module, int32))(m, v115)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L9
	} else {
		goto L43
	}
L43:
	;
	if v127 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	if v127 != 0 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	goto L46
L46:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v115)+12))
	m.T0[v156].(func(*base.Module, int32, int32, int32))(m, v115, l0+int32(2), int32(8))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L9
	} else {
		goto L51
	}
L47:
	;
	base.MemoryFill(m, v19+int32(16), int32(0), v127)
	goto L49
L48:
	;
	goto L49
L49:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v115)+12))
	m.T0[v152].(func(*base.Module, int32, int32, int32))(m, v115, v19+int32(16), v127)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L9
	} else {
		goto L50
	}
L50:
	;
	goto L46
L51:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v115)+12))
	m.T0[v159].(func(*base.Module, int32, int32, int32))(m, v115, l2, l3)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L9
	} else {
		goto L52
	}
L52:
	;
	v163 = v19 + int32(16)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v115)+16))
	m.T0[v164].(func(*base.Module, int32, int32))(m, v115, v163)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L9
	} else {
		goto L53
	}
L53:
	;
	if base.Ui32(v128) <= base.Ui32(v117) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	if v128 == int32(0) {
		goto L14
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	if v117 != 0 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	base.MemoryCopy(m, v135, v163, v128)
	goto L14
L58:
	;
	base.MemoryCopy(m, v135, v19+int32(16), v117)
	goto L60
L59:
	;
	goto L60
L60:
	;
	v127 = v127 + int32(1)
	v128 = v128 - v117
	v135 = v117 + v135
	goto L41
L61:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+43)))
	if v183 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v349 = int32(0)
	goto L104
L63:
	;
	v186 = int32(8)
	v189 = int32(16)
	v195 = (v178&int32(15) | v189) << (uint(int32(base.Ui32(v178)>>(uint(int32(4))%32))+int32(6)) % 32)
	v197 = l3 + v186
	if base.Ui32(l3+v189) < base.Ui32(v195) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v202 = v186
	goto L66
L65:
	;
	v202 = v195 - v197
	goto L66
L66:
	;
	v203 = v202 + v197
	v204 = v203 + l3
	v206 = l0 + int32(2)
	v219 = v183
	v220 = v5
	v221 = l0 + int32(11)
	goto L67
L67:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v179)+8))
	m.T0[v225].(func(*base.Module, int32))(m, v179)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L9
	} else {
		goto L69
	}
L69:
	;
	if v220 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	if v220 != 0 {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	goto L72
L72:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v179)+12))
	m.T0[v238].(func(*base.Module, int32, int32, int32))(m, v179, v206, int32(8))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L9
	} else {
		goto L77
	}
L73:
	;
	base.MemoryFill(m, v19+int32(16), int32(0), v220)
	goto L75
L74:
	;
	goto L75
L75:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v179)+12))
	m.T0[v234].(func(*base.Module, int32, int32, int32))(m, v179, v19+int32(16), v220)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L9
	} else {
		goto L76
	}
L76:
	;
	goto L72
L77:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v179)+12))
	m.T0[v241].(func(*base.Module, int32, int32, int32))(m, v179, l2, l3)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L9
	} else {
		goto L78
	}
L78:
	;
	if base.Ui32(v195) <= base.Ui32(v197) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v318 = v19 + int32(16)
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v179)+16))
	m.T0[v319].(func(*base.Module, int32, int32))(m, v179, v318)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L9
	} else {
		goto L95
	}
L80:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v179)+12))
	m.T0[v245].(func(*base.Module, int32, int32, int32))(m, v179, v206, v202)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L9
	} else {
		goto L81
	}
L81:
	;
	if base.Ui32(v204) < base.Ui32(v195) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v249 = v204
	goto L85
L83:
	;
	v281 = v203
	goto L84
L84:
	;
	if base.Ui32(v195) <= base.Ui32(v281) {
		goto L79
	} else {
		goto L93
	}
L85:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v179)+12))
	m.T0[v265].(func(*base.Module, int32, int32, int32))(m, v179, l2, l3)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L9
	} else {
		goto L87
	}
L86:
	;
	v281 = v277
	goto L84
L87:
	;
	v268 = int32(8)
	if base.Ui32(v249+v268) < base.Ui32(v195) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v273 = v268
	goto L90
L89:
	;
	v273 = v195 - v249
	goto L90
L90:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v179)+12))
	m.T0[v274].(func(*base.Module, int32, int32, int32))(m, v179, v206, v273)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L9
	} else {
		goto L91
	}
L91:
	;
	v277 = v249 + v273
	v278 = v277 + l3
	if base.Ui32(v278) < base.Ui32(v195) {
		v249 = v278
		goto L85
	} else {
		goto L92
	}
L92:
	;
	goto L86
L93:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v179)+12))
	m.T0[v298].(func(*base.Module, int32, int32, int32))(m, v179, l2, v195-v281)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L9
	} else {
		goto L94
	}
L94:
	;
	goto L79
L95:
	;
	if base.Ui32(v219) <= base.Ui32(v181) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	if v219 == int32(0) {
		goto L62
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	if v181 != 0 {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	base.MemoryCopy(m, v221, v318, v219)
	goto L62
L100:
	;
	base.MemoryCopy(m, v221, v19+int32(16), v181)
	goto L102
L101:
	;
	goto L102
L102:
	;
	v219 = v219 - v181
	v220 = v220 + int32(1)
	v221 = v181 + v221
	goto L67
L103:
	;
	v400 = v349
	goto L12
L104:
	;
	base.MemoryFill(m, v19+int32(16), v349, int32(64))
	goto L106
L106:
	;
	goto L103
L107:
	;
	v400 = v371
	goto L12
L108:
	;
	base.MemoryFill(m, v19+int32(16), v371, int32(64))
	goto L110
L110:
	;
	goto L107
L111:
	;
	v400 = v393
	goto L12
L112:
	;
	base.MemoryFill(m, v19+int32(16), v393, int32(64))
	goto L114
L114:
	;
	goto L111
L115:
	;
	v420 = v400
	goto L5
}
func F_pgp_set_s2k_digest_algo(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	v6 = F_pg_strcasecmp(m, int32(_a_F_pgp_set_s2k_digest_algo_0), l1)
	mBase = m.M
	if v6 == int32(0) {
		v39 = int32(_a_F_pgp_set_s2k_digest_algo_1)
		v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
		v41 = v40
	} else {
		v11 = F_pg_strcasecmp(m, int32(_a_F_pgp_set_s2k_digest_algo_2), l1)
		mBase = m.M
		if v11 == int32(0) {
			v39 = int32(_a_F_pgp_set_s2k_digest_algo_3)
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
			v41 = v40
		} else {
			v16 = F_pg_strcasecmp(m, int32(_a_F_pgp_set_s2k_digest_algo_4), l1)
			mBase = m.M
			if v16 == int32(0) {
				v39 = int32(_a_F_pgp_set_s2k_digest_algo_5)
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
				v41 = v40
			} else {
				v21 = F_pg_strcasecmp(m, int32(_a_F_pgp_set_s2k_digest_algo_6), l1)
				mBase = m.M
				if v21 == int32(0) {
					v39 = int32(_a_F_pgp_set_s2k_digest_algo_7)
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
					v41 = v40
				} else {
					v26 = F_pg_strcasecmp(m, int32(_a_F_pgp_set_s2k_digest_algo_8), l1)
					mBase = m.M
					if v26 == int32(0) {
						v39 = int32(_a_F_pgp_set_s2k_digest_algo_9)
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
						v41 = v40
					} else {
						v31 = F_pg_strcasecmp(m, int32(_a_F_pgp_set_s2k_digest_algo_10), l1)
						mBase = m.M
						if v31 == int32(0) {
							v39 = int32(_a_F_pgp_set_s2k_digest_algo_11)
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
							v41 = v40
						} else {
							v36 = F_pg_strcasecmp(m, int32(_a_F_pgp_set_s2k_digest_algo_12), l1)
							mBase = m.M
							if v36 != 0 {
								v41 = int32(-104)
							} else {
								v39 = int32(_a_F_pgp_set_s2k_digest_algo_13)
								v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
								v41 = v40
							}
						}
					}
				}
			}
		}
	}
	if v41 < int32(0) {
		return v41
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v41
		return int32(0)
	}
}
func F_pgp_set_text_mode(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = l1
	return int32(0)
}
