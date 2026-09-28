package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_match_network_subset(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int64
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int64
	_ = v65
	var v66 int32
	_ = v66
	var v68 int64
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	v5 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v12 != int32(7) {
		v81 = v5
		m.G0 = v10 + int32(16)
		return v81
	} else {
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+32)))
		if v15 != 0 {
			v81 = v5
			m.G0 = v10 + int32(16)
			return v81
		} else {
			v16 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
			v17 = int32(869)
			if l2 != 0 {
				v21 = int32(4)
			} else {
				v21 = int32(5)
			}
			v22 = F_get_opfamily_member_for_cmptype(m, l3, v17, v17, v21)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				if v22 == int32(0) {
					v81 = v5
					m.G0 = v10 + int32(16)
					return v81
				} else {
					v29 = int32(-1)
					v30 = int32(0)
					v34 = F_DirectFunctionCall1Coll(m, int32(1593), v30, v16)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						v36 = int32(0)
						v38 = F_makeConst(m, int32(869), v29, v30, v29, v34, v36, v36)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							v41 = F_make_opclause(m, v22, l0, v38, int32(0))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v41
								*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v41
								v48 = F_list_make1_impl(m, int32(1), v10+int32(8))
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int32(0)
								} else {
									v50 = int32(869)
									v53 = F_get_opfamily_member_for_cmptype(m, l3, v50, v50, int32(2))
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int32(0)
									} else {
										if v53 == int32(0) {
											v81 = v5
											m.G0 = v10 + int32(16)
											return v81
										} else {
											v58 = int32(-1)
											v59 = int32(0)
											v65 = F_DirectFunctionCall1Coll(m, int32(1595), v59, v16)
											mBase = m.M
											v66 = m.ExcPending
											if v66 != 0 {
												return int32(0)
											} else {
												v68 = F_DirectFunctionCall2Coll(m, int32(1594), v59, v65, int64(-1))
												mBase = m.M
												v69 = m.ExcPending
												if v69 != 0 {
													return int32(0)
												} else {
													v70 = int32(0)
													v72 = F_makeConst(m, int32(869), v58, v59, v58, v68, v70, v70)
													mBase = m.M
													v73 = m.ExcPending
													if v73 != 0 {
														return int32(0)
													} else {
														v75 = F_make_opclause(m, v53, l0, v72, int32(0))
														mBase = m.M
														v76 = m.ExcPending
														if v76 != 0 {
															return int32(0)
														} else {
															v77 = F_lappend(m, v48, v75)
															mBase = m.M
															v78 = m.ExcPending
															if v78 != 0 {
																return int32(0)
															} else {
																v81 = v77
																m.G0 = v10 + int32(16)
																return v81
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
			}
		}
	}
}
func F_network_abbrev_convert(m *base.Module, l0 int64, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v37 int64
	_ = v37
	var v38 int64
	_ = v38
	var v40 int64
	_ = v40
	var v42 int64
	_ = v42
	var v45 int64
	_ = v45
	var v47 int64
	_ = v47
	var v49 int64
	_ = v49
	var v51 int64
	_ = v51
	var v75 int64
	_ = v75
	var v77 int64
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v88 int64
	_ = v88
	var v90 int64
	_ = v90
	var v95 int64
	_ = v95
	var v97 int64
	_ = v97
	var v98 int64
	_ = v98
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v121 int64
	_ = v121
	var v122 int64
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	v6 = int64(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v11 = F_pg_detoast_datum_packed(m, base.I32_wrap_i64(l0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		v15 = int32(1)
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
		if v17&v15 != 0 {
			v20 = v15
		} else {
			v20 = int32(4)
		}
		v21 = v11 + v20
		v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
		if v22 == int32(2) {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v21)+2))
			v26 = int32(16711935)
			v75 = base.I64_extend_i32_u(base.I32_rotr(v25&v26, int32(8)) | base.I32_rotr(v25, int32(24))&v26)
			v77 = v6
			v78 = int32(32)
		} else {
			v37 = *(*int64)(unsafe.Add(mBase, uint32(v21)+2))
			v38 = int64(56)
			v40 = int64(65280)
			v42 = int64(40)
			v45 = int64(16711680)
			v47 = int64(24)
			v49 = int64(4278190080)
			v51 = int64(8)
			v75 = v37<<(uint(v38)%64) | v37&v40<<(uint(v42)%64) | (v37&v45<<(uint(v47)%64) | v37&v49<<(uint(v51)%64)) | (int64(base.Ui64(v37)>>(uint(v51)%64))&v49 | int64(base.Ui64(v37)>>(uint(v47)%64))&v45 | (int64(base.Ui64(v37)>>(uint(v42)%64))&v40 | int64(base.Ui64(v37)>>(uint(v38)%64))))
			v77 = int64(-9223372036854775807 - 1)
			v78 = int32(0)
		}
		v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
		v82 = (v78 - v79) & int32(63)
		if v79 == int32(0) {
			v95 = v6
			v97 = v75
		} else {
			if base.Ui32(int32(63)) < base.Ui32(v79) {
				v95 = v75
				v97 = int64(0)
			} else {
				v88 = int64(-1)
				v90 = v88 << (uint(base.I64_extend_i32_u(v82)) % 64)
				v95 = v90 & v75
				v97 = v75 & (v90 ^ v88)
			}
		}
		v98 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
		*(*int64)(unsafe.Add(mBase, uint32(v9))) = v98 + int64(1)
		if v22 == int32(2) {
			v111 = v82 - int32(25)
			if base.Ui32(v111) <= base.Ui32(v82) {
				v114 = v111
			} else {
				v114 = int32(0)
			}
			v121 = v95<<(uint(int64(31))%64) | base.I64_extend_i32_u(v79)<<(uint(int64(25))%64) | int64(base.Ui64(v97)>>(uint(base.I64_extend_i32_u(v114))%64))
		} else {
			v121 = int64(base.Ui64(v95) >> (uint(int64(1)) % 64))
		}
		v122 = v77 | v121
		v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+8)))
		if v123 == int32(1) {
			v126 = int32(16)
			v127 = v9 + v126
			v136 = int32(711645284)
			v139 = base.I32_wrap_i64(int64(base.Ui64(v122)>>(uint(int64(32))%64))^v121) - int32(1636608428) ^ v136 - int32(1455628627)
			v144 = v139 ^ int32(-1636608428) - base.I32_rotl(v139, int32(25))
			v149 = v144 ^ v136 - base.I32_rotl(v144, v126)
			v153 = v149 ^ v139 - base.I32_rotl(v149, int32(4))
			v157 = v153 ^ v144 - base.I32_rotl(v153, int32(14))
			v161 = v157 ^ v149 - base.I32_rotl(v157, int32(24))
			v164 = *(*int32)(unsafe.Add(mBase, uint32(v127)+16))
			v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
			v167 = int32(32) - v166
			v169 = v164 + int32(base.Ui32(v161)>>(uint(v167)%32))
			v170 = v161 << (uint(v166) % 32)
			if v170 != 0 {
				v177 = int32(32) - (base.I32_clz(v170) ^ int32(31))
				v178 = int32(255)
				if base.Ui32(v167&v178) < base.Ui32(v177&v178) {
					v183 = v167 + int32(1)
				} else {
					v183 = v177
				}
				v187 = v183
			} else {
				v187 = v167 + int32(1)
			}
			v189 = v187 & int32(255)
			v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169))))
			if base.Ui32(v190) < base.Ui32(v189) {
				v192 = v189
			} else {
				v192 = v190
			}
			*(*uint8)(unsafe.Add(mBase, uint32(v169))) = uint8(v192)
		} else {
		}
		return v122
	}
}
func F_network_hostmask(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v13 = F_palloc0(m, int32(22))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int64(0)
		} else {
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
			v16 = int32(1)
			v17 = v15 & v16
			v19 = int32(4)
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
			v26 = v24 & v16
			if v26 != 0 {
				v27 = v16
			} else {
				v27 = v19
			}
			v28 = v8 + v27
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
			v31 = base.B2i32(v29 == int32(2))
			if v29 == int32(2) {
				v32 = int32(32)
			} else {
				v32 = int32(128)
			}
			v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)))
			v34 = v32 - v33
			if v34 != 0 {
				if v17 != 0 {
					v37 = int32(1)
				} else {
					v37 = int32(4)
				}
				if v29 == int32(2) {
					v43 = int32(3)
				} else {
					v43 = int32(15)
				}
				v44 = v34
				v45 = v43
				for {
					if int32(7) < v44 {
						v58 = int32(-1)
					} else {
						v58 = int32(base.Ui32(int32(255)) >> (uint(int32(8)-v44) % 32))
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v45+(v13+v37+int32(2))))) = uint8(v58)
					v62 = int32(8)
					if v44 <= v62 {
						v65 = v62
					} else {
						v65 = v44
					}
					v67 = v65 - int32(8)
					if v67 != 0 {
						v44 = v67
						v45 = v45 - int32(1)
						continue
					} else {
						break
					}
					break
				}
				v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
				v69 = int32(1)
				v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
				v78 = v68 & v69
				v80 = v71 & v69
			} else {
				v78 = v26
				v80 = v17
			}
			if v80 != 0 {
				v81 = v16
			} else {
				v81 = v19
			}
			v82 = v13 + v81
			if v78 != 0 {
				v85 = int32(1)
			} else {
				v85 = int32(4)
			}
			v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+v85))))
			*(*uint8)(unsafe.Add(mBase, uint32(v82))) = uint8(v87)
			v91 = int32(1)
			v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
			if v93&v91 != 0 {
				v96 = v91
			} else {
				v96 = int32(4)
			}
			v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+v96))))
			if v98 == int32(2) {
				v101 = int32(32)
			} else {
				v101 = int32(-128)
			}
			*(*uint8)(unsafe.Add(mBase, uint32(v82)+1)) = uint8(v101)
			if v87 == int32(2) {
				v107 = int32(40)
			} else {
				v107 = int32(88)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v13))) = v107
			return base.I64_extend_i32_u(v13)
		}
	}
}
func F_network_lt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v156 int32
	_ = v156
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v8 = F_pg_detoast_datum_packed(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int64(0)
		} else {
			v17 = int32(1)
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
			if v19&v17 != 0 {
				v22 = v17
			} else {
				v22 = int32(4)
			}
			v23 = v3 + v22
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
			v25 = int32(1)
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
			if v27&v25 != 0 {
				v30 = v25
			} else {
				v30 = int32(4)
			}
			v31 = v8 + v30
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
			if v24 == v32 {
				v34 = int32(2)
				v35 = v23 + v34
				v37 = v31 + v34
				v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
				if base.Ui32(v38) < base.Ui32(v39) {
					v41 = v23
				} else {
					v41 = v31
				}
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
				v44 = int32(base.Ui32(v42) >> (uint(int32(3)) % 32))
				v45 = F_memcmp(m, v35, v37, v44)
				mBase = m.M
				if v45 != 0 {
					v137 = v45
					v156 = v137
				} else {
					v47 = v42 & int32(7)
					if v47 == int32(0) {
						v128 = v38 - v39
						if v128 != 0 {
							v137 = v128
							v156 = v137
						} else {
							if v24 == int32(2) {
								v133 = int32(4)
							} else {
								v133 = int32(16)
							}
							v134 = F_memcmp(m, v35, v37, v133)
							mBase = m.M
							v156 = v134
						}
					} else {
						v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+v35))))
						v52 = int32(128)
						v53 = v51 & v52
						v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+v37))))
						if v53 != v55&v52 {
							v144 = v53
							if v144 != 0 {
								v147 = int32(1)
							} else {
								v147 = int32(-1)
							}
							v156 = v147
						} else {
							if v47 == int32(1) {
								v128 = v38 - v39
								if v128 != 0 {
									v137 = v128
									v156 = v137
								} else {
									if v24 == int32(2) {
										v133 = int32(4)
									} else {
										v133 = int32(16)
									}
									v134 = F_memcmp(m, v35, v37, v133)
									mBase = m.M
									v156 = v134
								}
							} else {
								v61 = int32(1)
								v63 = int32(128)
								v64 = v51 << (uint(v61) % 32) & v63
								if v64 != v55<<(uint(v61)%32)&v63 {
									v144 = v64
									if v144 != 0 {
										v147 = int32(1)
									} else {
										v147 = int32(-1)
									}
									v156 = v147
								} else {
									if base.Ui32(v47) < base.Ui32(int32(3)) {
										v128 = v38 - v39
										if v128 != 0 {
											v137 = v128
											v156 = v137
										} else {
											if v24 == int32(2) {
												v133 = int32(4)
											} else {
												v133 = int32(16)
											}
											v134 = F_memcmp(m, v35, v37, v133)
											mBase = m.M
											v156 = v134
										}
									} else {
										v72 = int32(2)
										v74 = int32(128)
										v75 = v51 << (uint(v72) % 32) & v74
										if v75 != v55<<(uint(v72)%32)&v74 {
											v144 = v75
											if v144 != 0 {
												v147 = int32(1)
											} else {
												v147 = int32(-1)
											}
											v156 = v147
										} else {
											if v47 == int32(3) {
												v128 = v38 - v39
												if v128 != 0 {
													v137 = v128
													v156 = v137
												} else {
													if v24 == int32(2) {
														v133 = int32(4)
													} else {
														v133 = int32(16)
													}
													v134 = F_memcmp(m, v35, v37, v133)
													mBase = m.M
													v156 = v134
												}
											} else {
												v83 = int32(3)
												v85 = int32(128)
												v86 = v51 << (uint(v83) % 32) & v85
												if v86 != v55<<(uint(v83)%32)&v85 {
													v144 = v86
													if v144 != 0 {
														v147 = int32(1)
													} else {
														v147 = int32(-1)
													}
													v156 = v147
												} else {
													if base.Ui32(v47) < base.Ui32(int32(5)) {
														v128 = v38 - v39
														if v128 != 0 {
															v137 = v128
															v156 = v137
														} else {
															if v24 == int32(2) {
																v133 = int32(4)
															} else {
																v133 = int32(16)
															}
															v134 = F_memcmp(m, v35, v37, v133)
															mBase = m.M
															v156 = v134
														}
													} else {
														v94 = int32(4)
														v96 = int32(128)
														v97 = v51 << (uint(v94) % 32) & v96
														if v97 != v55<<(uint(v94)%32)&v96 {
															v144 = v97
															if v144 != 0 {
																v147 = int32(1)
															} else {
																v147 = int32(-1)
															}
															v156 = v147
														} else {
															if v47 == int32(5) {
																v128 = v38 - v39
																if v128 != 0 {
																	v137 = v128
																	v156 = v137
																} else {
																	if v24 == int32(2) {
																		v133 = int32(4)
																	} else {
																		v133 = int32(16)
																	}
																	v134 = F_memcmp(m, v35, v37, v133)
																	mBase = m.M
																	v156 = v134
																}
															} else {
																v105 = int32(5)
																v107 = int32(128)
																v108 = v51 << (uint(v105) % 32) & v107
																if v108 != v55<<(uint(v105)%32)&v107 {
																	v144 = v108
																	if v144 != 0 {
																		v147 = int32(1)
																	} else {
																		v147 = int32(-1)
																	}
																	v156 = v147
																} else {
																	if v47 != int32(7) {
																		v128 = v38 - v39
																		if v128 != 0 {
																			v137 = v128
																			v156 = v137
																		} else {
																			if v24 == int32(2) {
																				v133 = int32(4)
																			} else {
																				v133 = int32(16)
																			}
																			v134 = F_memcmp(m, v35, v37, v133)
																			mBase = m.M
																			v156 = v134
																		}
																	} else {
																		v116 = int32(6)
																		v118 = int32(128)
																		v119 = v51 << (uint(v116) % 32) & v118
																		if v119 != v55<<(uint(v116)%32)&v118 {
																			v144 = v119
																			if v144 != 0 {
																				v147 = int32(1)
																			} else {
																				v147 = int32(-1)
																			}
																			v156 = v147
																		} else {
																			v128 = v38 - v39
																			if v128 != 0 {
																				v137 = v128
																				v156 = v137
																			} else {
																				if v24 == int32(2) {
																					v133 = int32(4)
																				} else {
																					v133 = int32(16)
																				}
																				v134 = F_memcmp(m, v35, v37, v133)
																				mBase = m.M
																				v156 = v134
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
							}
						}
					}
				}
			} else {
				v137 = v24 - v32
				v156 = v137
			}
			return base.I64_extend_i32_u(int32(base.Ui32(v156) >> (uint(int32(31)) % 32)))
		}
	}
}
func F_network_out(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	v5 = m.G0
	v7 = v5 - int32(80)
	m.G0 = v7
	v9 = int32(1)
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v11&v9 != 0 {
		v14 = v9
	} else {
		v14 = int32(4)
	}
	v15 = l0 + v14
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
	v21 = v7 + int32(16)
	v22 = F_pg_inet_net_ntop(m, v16, v15+int32(2), v19, v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return int32(0)
	} else {
		if v22 != 0 {
			if l1 == int32(0) {
				v55 = F_pstrdup(m, v7+int32(16))
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int32(0)
				} else {
					m.G0 = v7 + int32(80)
					return v55
				}
			} else {
				v28 = int32(47)
				v29 = F___strchrnul(m, v21, v28)
				mBase = m.M
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
				if v31 == v28 {
					v35 = v29
				} else {
					v35 = int32(0)
				}
				if v35 != 0 {
					v55 = F_pstrdup(m, v7+int32(16))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						m.G0 = v7 + int32(80)
						return v55
					}
				} else {
					v36 = F_strlen(m, v21)
					mBase = m.M
					v37 = int32(1)
					v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
					if v39&v37 != 0 {
						v42 = v37
					} else {
						v42 = int32(4)
					}
					v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v42)+1)))
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v44
					v50 = F_pg_snprintf(m, v36+v21, int32(50)-v36, int32(_a_F_network_out_0), v7)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						v55 = F_pstrdup(m, v7+int32(16))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							m.G0 = v7 + int32(80)
							return v55
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50462850))
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_network_out_1), int32(0))
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_network_out_2), int32(150), int32(_a_F_network_out_3))
						mBase = m.M
						v76 = m.ExcPending
						if v76 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	}
}
func F_network_recv(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v14 = F_palloc0(m, int32(22))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v18 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v20 = int32(1)
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v22&v20 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v25 = v20
	goto L6
L5:
	;
	v25 = int32(4)
	goto L6
L6:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v14+v25))) = uint8(v18)
	if v18&int32(254) == int32(2) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	if v89 != 0 {
		goto L80
	} else {
		goto L81
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L73
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L66
	}
L10:
	;
	v32 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L59
	}
L13:
	;
	if v32 < int32(0) {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	v38 = int32(1)
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v40&v38 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v43 = v38
	goto L17
L16:
	;
	v43 = int32(4)
	goto L17
L17:
	;
	v44 = v14 + v43
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if v45 == int32(2) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v48 = int32(32)
	goto L20
L19:
	;
	v48 = int32(128)
	goto L20
L20:
	;
	if base.Ui32(v48) < base.Ui32(v32) {
		goto L9
	} else {
		goto L21
	}
L21:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+1)) = uint8(v32)
	v51 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v53 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v55 = int32(4)
	v57 = int32(1)
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v59&v57 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v62 = v57
	goto L26
L25:
	;
	v62 = v55
	goto L26
L26:
	;
	v63 = v14 + v62
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	if v64 == int32(2) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v67 = v55
	goto L29
L28:
	;
	v67 = int32(16)
	goto L29
L29:
	;
	if v53 != v67 {
		goto L8
	} else {
		goto L30
	}
L30:
	;
	v74 = int32(0)
	goto L31
L31:
	;
	v81 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L33
	}
L32:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	v89 = v87 & int32(1)
	if l1 == int32(0) {
		goto L7
	} else {
		goto L35
	}
L33:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v74+(v63+int32(2))))) = uint8(v81)
	v85 = v74 + int32(1)
	if v85 != v53 {
		v74 = v85
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	if v89 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v96 = int32(1)
	goto L38
L37:
	;
	v96 = int32(4)
	goto L38
L38:
	;
	v97 = v14 + v96
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	v100 = base.B2i32(v98 == int32(2))
	if v98 == int32(2) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v101 = int32(32)
	goto L41
L40:
	;
	v101 = int32(128)
	goto L41
L41:
	;
	if v32 == v101 {
		goto L7
	} else {
		goto L42
	}
L42:
	;
	v104 = int32(base.Ui32(v32) >> (uint(int32(3)) % 32))
	if v98 == int32(2) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v107 = int32(4)
	goto L45
L44:
	;
	v107 = int32(16)
	goto L45
L45:
	;
	if base.Ui32(v107) <= base.Ui32(v104) {
		goto L7
	} else {
		goto L46
	}
L46:
	;
	v110 = v97 + int32(2)
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110+v104))))
	if v112<<(uint(v32&int32(7))%32)&int32(255) != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L54
	}
L48:
	;
	v119 = v104 + int32(1)
	if v119 == v107 {
		goto L7
	} else {
		goto L49
	}
L49:
	;
	v123 = v119
	goto L50
L50:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123+v110))))
	if v130 != 0 {
		goto L47
	} else {
		goto L52
	}
L51:
	;
	goto L7
L52:
	;
	v132 = v123 + int32(1)
	if v107 != v132 {
		v123 = v132
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	F_errmsg(m, int32(_a_F_network_recv_0), int32(0))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v155 = F_errdetail(m, int32(_a_F_network_recv_1), int32(0))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(_a_F_network_recv_2), int32(239), int32(_a_F_network_recv_3))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L59:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	if l1 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v171 = int32(_a_F_network_recv_4)
	goto L63
L62:
	;
	v171 = int32(_a_F_network_recv_5)
	goto L63
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v171
	F_errmsg(m, int32(_a_F_network_recv_6), v11+int32(32))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_network_recv_2), int32(208), int32(_a_F_network_recv_3))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	if l1 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v193 = int32(_a_F_network_recv_4)
	goto L70
L69:
	;
	v193 = int32(_a_F_network_recv_5)
	goto L70
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v193
	F_errmsg(m, int32(_a_F_network_recv_7), v11)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_network_recv_2), int32(215), int32(_a_F_network_recv_3))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L73:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	if l1 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v212 = int32(_a_F_network_recv_4)
	goto L77
L76:
	;
	v212 = int32(_a_F_network_recv_5)
	goto L77
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v212
	F_errmsg(m, int32(_a_F_network_recv_8), v11+int32(16))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_network_recv_2), int32(224), int32(_a_F_network_recv_3))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L80:
	;
	v236 = int32(1)
	goto L82
L81:
	;
	v236 = int32(4)
	goto L82
L82:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+v236))))
	if v238 == int32(2) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v241 = int32(40)
	goto L85
L84:
	;
	v241 = int32(88)
	goto L85
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v241
	m.G0 = v11 + int32(48)
	return v14
}
func F_network_sup(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v155 int64
	_ = v155
	v5 = int64(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v14 = int32(1)
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v16&v14 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int64(1)
L5:
	;
	return v155
L6:
	;
	v19 = v14
	goto L8
L7:
	;
	v19 = int32(4)
	goto L8
L8:
	;
	v20 = v7 + v19
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	v22 = int32(1)
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v24&v22 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v27 = v22
	goto L11
L10:
	;
	v27 = int32(4)
	goto L11
L11:
	;
	v28 = v12 + v27
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if v21 != v29 {
		v155 = v5
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)))
	if base.Ui32(v32) <= base.Ui32(v31) {
		v155 = v5
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v34 = int32(2)
	v35 = v20 + v34
	v37 = v28 + v34
	v39 = int32(base.Ui32(v31) >> (uint(int32(3)) % 32))
	if base.Ui32(int32(4)) <= base.Ui32(v39) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	if v101 != 0 {
		v155 = v5
		goto L5
	} else {
		goto L32
	}
L15:
	;
	v101 = int32(0)
	goto L14
L16:
	;
	v75 = v70
	v76 = v71
	v77 = v72
	goto L26
L17:
	;
	if (v35|v37)&int32(3) != 0 {
		v70 = v35
		v71 = v37
		v72 = v39
		goto L16
	} else {
		goto L20
	}
L18:
	;
	v63 = v35
	v64 = v37
	v65 = v39
	goto L19
L19:
	;
	if v65 == int32(0) {
		goto L15
	} else {
		goto L25
	}
L20:
	;
	v47 = v35
	v48 = v37
	v49 = v39
	goto L21
L21:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	if v52 != v53 {
		v70 = v47
		v71 = v48
		v72 = v49
		goto L16
	} else {
		goto L23
	}
L22:
	;
	v63 = v58
	v64 = v56
	v65 = v60
	goto L19
L23:
	;
	v55 = int32(4)
	v56 = v48 + v55
	v58 = v47 + v55
	v60 = v49 - v55
	if base.Ui32(int32(3)) < base.Ui32(v60) {
		v47 = v58
		v48 = v56
		v49 = v60
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v70 = v63
	v71 = v64
	v72 = v65
	goto L16
L26:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	if v80 == v81 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v101 = v80 - v81
	goto L14
L28:
	;
	v83 = int32(1)
	v88 = v77 - v83
	if v88 != 0 {
		v75 = v75 + v83
		v76 = v76 + v83
		v77 = v88
		goto L26
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	goto L27
L31:
	;
	goto L15
L32:
	;
	v103 = v31 & int32(7)
	if v103 == int32(0) {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35+v39))))
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+v39))))
	v110 = v107 ^ v109
	if base.Ui32(int32(127)) < base.Ui32(v110) {
		v155 = v5
		goto L5
	} else {
		goto L34
	}
L34:
	;
	if v103 == int32(1) {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	if v110<<(uint(int32(1))%32)&int32(128) != 0 {
		v155 = v5
		goto L5
	} else {
		goto L36
	}
L36:
	;
	if base.Ui32(v103) < base.Ui32(int32(3)) {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	if v110<<(uint(int32(2))%32)&int32(128) != 0 {
		v155 = v5
		goto L5
	} else {
		goto L38
	}
L38:
	;
	if v103 == int32(3) {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	if v110<<(uint(int32(3))%32)&int32(128) != 0 {
		v155 = v5
		goto L5
	} else {
		goto L40
	}
L40:
	;
	if base.Ui32(v103) < base.Ui32(int32(5)) {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	if v110<<(uint(int32(4))%32)&int32(128) != 0 {
		v155 = v5
		goto L5
	} else {
		goto L42
	}
L42:
	;
	if v103 == int32(5) {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	if v110<<(uint(int32(5))%32)&int32(128) != 0 {
		v155 = v5
		goto L5
	} else {
		goto L44
	}
L44:
	;
	if v103 != int32(7) {
		v155 = int64(1)
		goto L5
	} else {
		goto L45
	}
L45:
	;
	v155 = base.I64_extend_i32_u(base.B2i32(v110&int32(2) == int32(0)))
	goto L5
}
func F_network_supeq(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v155 int64
	_ = v155
	v5 = int64(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v14 = int32(1)
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v16&v14 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int64(1)
L5:
	;
	return v155
L6:
	;
	v19 = v14
	goto L8
L7:
	;
	v19 = int32(4)
	goto L8
L8:
	;
	v20 = v7 + v19
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	v22 = int32(1)
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v24&v22 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v27 = v22
	goto L11
L10:
	;
	v27 = int32(4)
	goto L11
L11:
	;
	v28 = v12 + v27
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if v21 != v29 {
		v155 = v5
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)))
	if base.Ui32(v32) < base.Ui32(v31) {
		v155 = v5
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v34 = int32(2)
	v35 = v20 + v34
	v37 = v28 + v34
	v39 = int32(base.Ui32(v31) >> (uint(int32(3)) % 32))
	if base.Ui32(int32(4)) <= base.Ui32(v39) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	if v101 != 0 {
		v155 = v5
		goto L5
	} else {
		goto L32
	}
L15:
	;
	v101 = int32(0)
	goto L14
L16:
	;
	v75 = v70
	v76 = v71
	v77 = v72
	goto L26
L17:
	;
	if (v35|v37)&int32(3) != 0 {
		v70 = v35
		v71 = v37
		v72 = v39
		goto L16
	} else {
		goto L20
	}
L18:
	;
	v63 = v35
	v64 = v37
	v65 = v39
	goto L19
L19:
	;
	if v65 == int32(0) {
		goto L15
	} else {
		goto L25
	}
L20:
	;
	v47 = v35
	v48 = v37
	v49 = v39
	goto L21
L21:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	if v52 != v53 {
		v70 = v47
		v71 = v48
		v72 = v49
		goto L16
	} else {
		goto L23
	}
L22:
	;
	v63 = v58
	v64 = v56
	v65 = v60
	goto L19
L23:
	;
	v55 = int32(4)
	v56 = v48 + v55
	v58 = v47 + v55
	v60 = v49 - v55
	if base.Ui32(int32(3)) < base.Ui32(v60) {
		v47 = v58
		v48 = v56
		v49 = v60
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v70 = v63
	v71 = v64
	v72 = v65
	goto L16
L26:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	if v80 == v81 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v101 = v80 - v81
	goto L14
L28:
	;
	v83 = int32(1)
	v88 = v77 - v83
	if v88 != 0 {
		v75 = v75 + v83
		v76 = v76 + v83
		v77 = v88
		goto L26
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	goto L27
L31:
	;
	goto L15
L32:
	;
	v103 = v31 & int32(7)
	if v103 == int32(0) {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35+v39))))
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+v39))))
	v110 = v107 ^ v109
	if base.Ui32(int32(127)) < base.Ui32(v110) {
		v155 = v5
		goto L5
	} else {
		goto L34
	}
L34:
	;
	if v103 == int32(1) {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	if v110<<(uint(int32(1))%32)&int32(128) != 0 {
		v155 = v5
		goto L5
	} else {
		goto L36
	}
L36:
	;
	if base.Ui32(v103) < base.Ui32(int32(3)) {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	if v110<<(uint(int32(2))%32)&int32(128) != 0 {
		v155 = v5
		goto L5
	} else {
		goto L38
	}
L38:
	;
	if v103 == int32(3) {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	if v110<<(uint(int32(3))%32)&int32(128) != 0 {
		v155 = v5
		goto L5
	} else {
		goto L40
	}
L40:
	;
	if base.Ui32(v103) < base.Ui32(int32(5)) {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	if v110<<(uint(int32(4))%32)&int32(128) != 0 {
		v155 = v5
		goto L5
	} else {
		goto L42
	}
L42:
	;
	if v103 == int32(5) {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	if v110<<(uint(int32(5))%32)&int32(128) != 0 {
		v155 = v5
		goto L5
	} else {
		goto L44
	}
L44:
	;
	if v103 != int32(7) {
		v155 = int64(1)
		goto L5
	} else {
		goto L45
	}
L45:
	;
	v155 = base.I64_extend_i32_u(base.B2i32(v110&int32(2) == int32(0)))
	goto L5
}
