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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
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
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
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
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	v5 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v11 != int32(7) {
		v81 = v5
		m.G0 = v9 + int32(16)
		return v81
	} else {
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
		if v14 != 0 {
			v81 = v5
			m.G0 = v9 + int32(16)
			return v81
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
			v16 = int32(869)
			if l2 != 0 {
				v20 = int32(4)
			} else {
				v20 = int32(5)
			}
			v21 = F_get_opfamily_member_for_cmptype(m, l3, v16, v16, v20)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				if v21 == int32(0) {
					v81 = v5
					m.G0 = v9 + int32(16)
					return v81
				} else {
					v28 = int32(-1)
					v29 = int32(0)
					v33 = F_DirectFunctionCall1Coll(m, int32(1462), v29, v15)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						v35 = int32(0)
						v37 = F_makeConst(m, int32(869), v28, v29, v28, v33, v35, v35)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							v40 = F_make_opclause(m, v21, l0, v37, int32(0))
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v40
								*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v40
								v47 = F_list_make1_impl(m, int32(1), v9+int32(8))
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return int32(0)
								} else {
									v49 = int32(869)
									v52 = F_get_opfamily_member_for_cmptype(m, l3, v49, v49, int32(2))
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return int32(0)
									} else {
										if v52 == int32(0) {
											v81 = v5
											m.G0 = v9 + int32(16)
											return v81
										} else {
											v57 = int32(-1)
											v58 = int32(0)
											v64 = F_DirectFunctionCall1Coll(m, int32(1464), v58, v15)
											mBase = m.M
											v65 = m.ExcPending
											if v65 != 0 {
												return int32(0)
											} else {
												v67 = F_DirectFunctionCall2Coll(m, int32(1463), v58, v64, int32(-1))
												mBase = m.M
												v68 = m.ExcPending
												if v68 != 0 {
													return int32(0)
												} else {
													v69 = int32(0)
													v71 = F_makeConst(m, int32(869), v57, v58, v57, v67, v69, v69)
													mBase = m.M
													v72 = m.ExcPending
													if v72 != 0 {
														return int32(0)
													} else {
														v74 = F_make_opclause(m, v52, l0, v71, int32(0))
														mBase = m.M
														v75 = m.ExcPending
														if v75 != 0 {
															return int32(0)
														} else {
															v76 = F_lappend(m, v47, v74)
															mBase = m.M
															v77 = m.ExcPending
															if v77 != 0 {
																return int32(0)
															} else {
																v81 = v76
																m.G0 = v9 + int32(16)
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
func F_network_abbrev_convert(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v51 int64
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v6 = F_pg_detoast_datum_packed(m, l0)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = int32(1)
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
		if v12&v10 != 0 {
			v15 = v10
		} else {
			v15 = int32(4)
		}
		v16 = v6 + v15
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
		v22 = int32(0)
		v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
		if v23 == v22 {
			v50 = v22
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v16)+2))
			v27 = int32(24)
			v29 = int32(65280)
			v31 = int32(8)
			v41 = v26<<(uint(v27)%32) | v26&v29<<(uint(v31)%32) | (int32(base.Ui32(v26)>>(uint(v31)%32))&v29 | int32(base.Ui32(v26)>>(uint(v27)%32)))
			if base.Ui32(int32(31)) < base.Ui32(v23) {
				v50 = v41
			} else {
				v50 = int32(-1) << (uint(int32(0)-v23) % 32) & v41
			}
		}
		v51 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
		*(*int64)(unsafe.Add(mBase, uint32(v5))) = v51 + int64(1)
		v55 = int32(1)
		v57 = int32(base.Ui32(v50)>>(uint(v55)%32)) | base.B2i32(v17 != int32(2))<<(uint(int32(31))%32)
		v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+8)))
		if v58 == v55 {
			v61 = int32(16)
			v62 = v5 + v61
			v67 = int32(711645284)
			v70 = v57 - int32(1636608428) ^ v67 - int32(1455628627)
			v75 = v70 ^ int32(-1636608428) - base.I32_rotl(v70, int32(25))
			v80 = v75 ^ v67 - base.I32_rotl(v75, v61)
			v84 = v80 ^ v70 - base.I32_rotl(v80, int32(4))
			v88 = v84 ^ v75 - base.I32_rotl(v84, int32(14))
			v92 = v88 ^ v80 - base.I32_rotl(v88, int32(24))
			v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
			v97 = int32(32) - v96
			v98 = v92 << (uint(v96) % 32)
			if v98 != 0 {
				v105 = int32(32) - (base.I32_clz(v98) ^ int32(31))
				v106 = int32(255)
				if base.Ui32(v97&v106) < base.Ui32(v105&v106) {
					v111 = v97 + int32(1)
				} else {
					v111 = v105
				}
				v115 = v111
			} else {
				v115 = v97 + int32(1)
			}
			v116 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
			v118 = v116 + int32(base.Ui32(v92)>>(uint(v97)%32))
			v120 = v115 & int32(255)
			v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
			if base.Ui32(v121) < base.Ui32(v120) {
				v123 = v120
			} else {
				v123 = v121
			}
			*(*uint8)(unsafe.Add(mBase, uint32(v118))) = uint8(v123)
		} else {
		}
		return v57
	}
}
func F_network_hostmask(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v15 = F_palloc0(m, int32(22))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = int32(4)
			v18 = v15 + v17
			v19 = int32(1)
			v20 = v15 + v19
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
			v23 = v21 & v19
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
			v30 = v28 & v19
			if v30 != 0 {
				v31 = v19
			} else {
				v31 = v17
			}
			v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+v31))))
			v35 = base.B2i32(v33 == int32(2))
			if v33 == int32(2) {
				v36 = int32(32)
			} else {
				v36 = int32(128)
			}
			if v30 != 0 {
				v41 = v10 + int32(1)
			} else {
				v41 = v10 + int32(4)
			}
			v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
			v43 = v36 - v42
			if v43 != 0 {
				if v23 != 0 {
					v44 = v20
				} else {
					v44 = v18
				}
				if v33 == int32(2) {
					v49 = int32(3)
				} else {
					v49 = int32(15)
				}
				v50 = v43
				v53 = v49
				for {
					if int32(7) < v50 {
						v66 = int32(-1)
					} else {
						v66 = int32(base.Ui32(int32(255)) >> (uint(int32(8)-v50) % 32))
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v53+(v44+int32(2))))) = uint8(v66)
					v70 = int32(8)
					if v50 <= v70 {
						v73 = v70
					} else {
						v73 = v50
					}
					v75 = v73 - int32(8)
					if v75 != 0 {
						v50 = v75
						v53 = v53 - int32(1)
						continue
					} else {
						break
					}
					break
				}
				v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
				v77 = int32(1)
				v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
				v85 = v79 & v77
				v86 = v76 & v77
			} else {
				v85 = v23
				v86 = v30
			}
			if v85 != 0 {
				v92 = int32(1)
			} else {
				v92 = int32(4)
			}
			if v86 != 0 {
				v96 = int32(1)
			} else {
				v96 = int32(4)
			}
			v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+v96))))
			*(*uint8)(unsafe.Add(mBase, uint32(v15+v92))) = uint8(v98)
			if v85 != 0 {
				v100 = v20
			} else {
				v100 = v18
			}
			v103 = int32(1)
			v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
			if v105&v103 != 0 {
				v108 = v103
			} else {
				v108 = int32(4)
			}
			v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+v108))))
			if v110 == int32(2) {
				v113 = int32(32)
			} else {
				v113 = int32(-128)
			}
			*(*uint8)(unsafe.Add(mBase, uint32(v100)+1)) = uint8(v113)
			v117 = int32(1)
			v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
			if v119&v117 != 0 {
				v122 = v117
			} else {
				v122 = int32(4)
			}
			v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+v122))))
			if v124 == int32(2) {
				v127 = int32(40)
			} else {
				v127 = int32(88)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v15))) = v127
			return v15
		}
	}
}
func F_network_lt(m *base.Module, l0 int32) int32 {
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v155 int32
	_ = v155
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v8 = F_pg_detoast_datum_packed(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
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
					v41 = v38
				} else {
					v41 = v39
				}
				v43 = int32(base.Ui32(v41) >> (uint(int32(3)) % 32))
				v44 = F_memcmp(m, v35, v37, v43)
				mBase = m.M
				if v44 != 0 {
					v136 = v44
					v155 = v136
				} else {
					v46 = v41 & int32(7)
					if v46 == int32(0) {
						v127 = v38 - v39
						if v127 != 0 {
							v136 = v127
							v155 = v136
						} else {
							if v24 == int32(2) {
								v132 = int32(4)
							} else {
								v132 = int32(16)
							}
							v133 = F_memcmp(m, v35, v37, v132)
							mBase = m.M
							v155 = v133
						}
					} else {
						v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+v35))))
						v51 = int32(128)
						v52 = v50 & v51
						v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+v37))))
						if v52 != v54&v51 {
							v143 = v52
							if v143 != 0 {
								v146 = int32(1)
							} else {
								v146 = int32(-1)
							}
							v155 = v146
						} else {
							if v46 == int32(1) {
								v127 = v38 - v39
								if v127 != 0 {
									v136 = v127
									v155 = v136
								} else {
									if v24 == int32(2) {
										v132 = int32(4)
									} else {
										v132 = int32(16)
									}
									v133 = F_memcmp(m, v35, v37, v132)
									mBase = m.M
									v155 = v133
								}
							} else {
								v60 = int32(1)
								v62 = int32(128)
								v63 = v50 << (uint(v60) % 32) & v62
								if v63 != v54<<(uint(v60)%32)&v62 {
									v143 = v63
									if v143 != 0 {
										v146 = int32(1)
									} else {
										v146 = int32(-1)
									}
									v155 = v146
								} else {
									if base.Ui32(v46) < base.Ui32(int32(3)) {
										v127 = v38 - v39
										if v127 != 0 {
											v136 = v127
											v155 = v136
										} else {
											if v24 == int32(2) {
												v132 = int32(4)
											} else {
												v132 = int32(16)
											}
											v133 = F_memcmp(m, v35, v37, v132)
											mBase = m.M
											v155 = v133
										}
									} else {
										v71 = int32(2)
										v73 = int32(128)
										v74 = v50 << (uint(v71) % 32) & v73
										if v74 != v54<<(uint(v71)%32)&v73 {
											v143 = v74
											if v143 != 0 {
												v146 = int32(1)
											} else {
												v146 = int32(-1)
											}
											v155 = v146
										} else {
											if v46 == int32(3) {
												v127 = v38 - v39
												if v127 != 0 {
													v136 = v127
													v155 = v136
												} else {
													if v24 == int32(2) {
														v132 = int32(4)
													} else {
														v132 = int32(16)
													}
													v133 = F_memcmp(m, v35, v37, v132)
													mBase = m.M
													v155 = v133
												}
											} else {
												v82 = int32(3)
												v84 = int32(128)
												v85 = v50 << (uint(v82) % 32) & v84
												if v85 != v54<<(uint(v82)%32)&v84 {
													v143 = v85
													if v143 != 0 {
														v146 = int32(1)
													} else {
														v146 = int32(-1)
													}
													v155 = v146
												} else {
													if base.Ui32(v46) < base.Ui32(int32(5)) {
														v127 = v38 - v39
														if v127 != 0 {
															v136 = v127
															v155 = v136
														} else {
															if v24 == int32(2) {
																v132 = int32(4)
															} else {
																v132 = int32(16)
															}
															v133 = F_memcmp(m, v35, v37, v132)
															mBase = m.M
															v155 = v133
														}
													} else {
														v93 = int32(4)
														v95 = int32(128)
														v96 = v50 << (uint(v93) % 32) & v95
														if v96 != v54<<(uint(v93)%32)&v95 {
															v143 = v96
															if v143 != 0 {
																v146 = int32(1)
															} else {
																v146 = int32(-1)
															}
															v155 = v146
														} else {
															if v46 == int32(5) {
																v127 = v38 - v39
																if v127 != 0 {
																	v136 = v127
																	v155 = v136
																} else {
																	if v24 == int32(2) {
																		v132 = int32(4)
																	} else {
																		v132 = int32(16)
																	}
																	v133 = F_memcmp(m, v35, v37, v132)
																	mBase = m.M
																	v155 = v133
																}
															} else {
																v104 = int32(5)
																v106 = int32(128)
																v107 = v50 << (uint(v104) % 32) & v106
																if v107 != v54<<(uint(v104)%32)&v106 {
																	v143 = v107
																	if v143 != 0 {
																		v146 = int32(1)
																	} else {
																		v146 = int32(-1)
																	}
																	v155 = v146
																} else {
																	if v46 != int32(7) {
																		v127 = v38 - v39
																		if v127 != 0 {
																			v136 = v127
																			v155 = v136
																		} else {
																			if v24 == int32(2) {
																				v132 = int32(4)
																			} else {
																				v132 = int32(16)
																			}
																			v133 = F_memcmp(m, v35, v37, v132)
																			mBase = m.M
																			v155 = v133
																		}
																	} else {
																		v115 = int32(6)
																		v117 = int32(128)
																		v118 = v50 << (uint(v115) % 32) & v117
																		if v118 != v54<<(uint(v115)%32)&v117 {
																			v143 = v118
																			if v143 != 0 {
																				v146 = int32(1)
																			} else {
																				v146 = int32(-1)
																			}
																			v155 = v146
																		} else {
																			v127 = v38 - v39
																			if v127 != 0 {
																				v136 = v127
																				v155 = v136
																			} else {
																				if v24 == int32(2) {
																					v132 = int32(4)
																				} else {
																					v132 = int32(16)
																				}
																				v133 = F_memcmp(m, v35, v37, v132)
																				mBase = m.M
																				v155 = v133
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
				v136 = v24 - v32
				v155 = v136
			}
			return int32(base.Ui32(v155) >> (uint(int32(31)) % 32))
		}
	}
}
func F_network_out(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	v7 = m.G0
	v9 = v7 - int32(80)
	m.G0 = v9
	v11 = int32(1)
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v15 = v13 & v11
	if v15 != 0 {
		v16 = v11
	} else {
		v16 = int32(4)
	}
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v16))))
	v20 = l0 + int32(1)
	v22 = l0 + int32(4)
	if v15 != 0 {
		v23 = v20
	} else {
		v23 = v22
	}
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	v29 = F_pg_inet_net_ntop(m, v18, v23+int32(2), v26, v9+int32(16))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		return int32(0)
	} else {
		if v29 != 0 {
			if l1 == int32(0) {
				v65 = F_pstrdup(m, v9+int32(16))
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return int32(0)
				} else {
					m.G0 = v9 + int32(80)
					return v65
				}
			} else {
				v37 = int32(47)
				v38 = F___strchrnul(m, v9+int32(16), v37)
				mBase = m.M
				v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
				if v40 == v37 {
					v44 = v38
				} else {
					v44 = int32(0)
				}
				if v44 != 0 {
					v65 = F_pstrdup(m, v9+int32(16))
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return int32(0)
					} else {
						m.G0 = v9 + int32(80)
						return v65
					}
				} else {
					v47 = F_strlen(m, v9+int32(16))
					mBase = m.M
					v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
					if v48&int32(1) != 0 {
						v51 = v20
					} else {
						v51 = v22
					}
					v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+1)))
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v52
					v60 = F_pg_snprintf(m, v47+(v9+int32(16)), int32(50)-v47, int32(40152), v9)
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						v65 = F_pstrdup(m, v9+int32(16))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							m.G0 = v9 + int32(80)
							return v65
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v74 = m.ExcPending
			if v74 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50462850))
				mBase = m.M
				v77 = m.ExcPending
				if v77 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(300456), int32(0))
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(509589), int32(152), int32(68574))
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
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
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
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
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	v16 = F_palloc0(m, int32(22))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = int32(1)
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v24&v22 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v27 = v22
	goto L6
L5:
	;
	v27 = int32(4)
	goto L6
L6:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v16+v27))) = uint8(v20)
	if v20&int32(254) == int32(2) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	if v99 != 0 {
		goto L86
	} else {
		goto L87
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L79
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L72
	}
L10:
	;
	v34 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
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
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L65
	}
L13:
	;
	if v34 < int32(0) {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	v40 = int32(1)
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v44 = v42 & v40
	if v44 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v45 = v40
	goto L17
L16:
	;
	v45 = int32(4)
	goto L17
L17:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+v45))))
	if v47 == int32(2) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v50 = int32(32)
	goto L20
L19:
	;
	v50 = int32(128)
	goto L20
L20:
	;
	if base.Ui32(v50) < base.Ui32(v34) {
		goto L9
	} else {
		goto L21
	}
L21:
	;
	v53 = v16 + int32(1)
	v55 = v16 + int32(4)
	if v44 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v56 = v53
	goto L24
L23:
	;
	v56 = v55
	goto L24
L24:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)) = uint8(v34)
	v58 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v60 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v62 = int32(4)
	v64 = int32(1)
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v68 = v66 & v64
	if v68 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v69 = v64
	goto L29
L28:
	;
	v69 = v62
	goto L29
L29:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+v69))))
	if v71 == int32(2) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v74 = v62
	goto L32
L31:
	;
	v74 = int32(16)
	goto L32
L32:
	;
	if v60 != v74 {
		goto L8
	} else {
		goto L33
	}
L33:
	;
	if v68 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v76 = v53
	goto L36
L35:
	;
	v76 = v55
	goto L36
L36:
	;
	v82 = int32(0)
	goto L37
L37:
	;
	v91 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L39
	}
L38:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v99 = v97 & int32(1)
	if l1 == int32(0) {
		goto L7
	} else {
		goto L41
	}
L39:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v82+(v76+int32(2))))) = uint8(v91)
	v95 = v82 + int32(1)
	if v95 != v60 {
		v82 = v95
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	if v99 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v104 = v53
	goto L44
L43:
	;
	v104 = v55
	goto L44
L44:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	v107 = base.B2i32(v105 == int32(2))
	if v105 == int32(2) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v108 = int32(32)
	goto L47
L46:
	;
	v108 = int32(128)
	goto L47
L47:
	;
	if v34 == v108 {
		goto L7
	} else {
		goto L48
	}
L48:
	;
	v111 = int32(base.Ui32(v34) >> (uint(int32(3)) % 32))
	if v105 == int32(2) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v114 = int32(4)
	goto L51
L50:
	;
	v114 = int32(16)
	goto L51
L51:
	;
	if base.Ui32(v114) <= base.Ui32(v111) {
		goto L7
	} else {
		goto L52
	}
L52:
	;
	v117 = v104 + int32(2)
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117+v111))))
	if v119<<(uint(v34&int32(7))%32)&int32(255) != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L60
	}
L54:
	;
	v126 = v111 + int32(1)
	if v126 == v114 {
		goto L7
	} else {
		goto L55
	}
L55:
	;
	v130 = v126
	goto L56
L56:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130+v117))))
	if v139 != 0 {
		goto L53
	} else {
		goto L58
	}
L57:
	;
	goto L7
L58:
	;
	v141 = v130 + int32(1)
	if v114 != v141 {
		v130 = v141
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	F_errmsg(m, int32(354935), int32(0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_errdetail(m, int32(636982), int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(509589), int32(241), int32(37421))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	if l1 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v182 = int32(233690)
	goto L69
L68:
	;
	v182 = int32(109653)
	goto L69
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v182
	F_errmsg(m, int32(354815), v13+int32(32))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(509589), int32(210), int32(37421))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L72:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	if l1 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v204 = int32(233690)
	goto L76
L75:
	;
	v204 = int32(109653)
	goto L76
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v204
	F_errmsg(m, int32(354861), v13)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(509589), int32(217), int32(37421))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L79:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	if l1 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v223 = int32(233690)
	goto L83
L82:
	;
	v223 = int32(109653)
	goto L83
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v223
	F_errmsg(m, int32(354897), v13+int32(16))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(509589), int32(226), int32(37421))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L86:
	;
	v249 = int32(1)
	goto L88
L87:
	;
	v249 = int32(4)
	goto L88
L88:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+v249))))
	if v251 == int32(2) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v254 = int32(40)
	goto L91
L90:
	;
	v254 = int32(88)
	goto L91
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v254
	m.G0 = v13 + int32(48)
	return v16
}
func F_network_sup(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
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
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v169 int32
	_ = v169
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v13 = int32(1)
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if v15&v13 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v18 = v13
	goto L6
L5:
	;
	v18 = int32(4)
	goto L6
L6:
	;
	v19 = v6 + v18
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	v21 = int32(1)
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v23&v21 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v26 = v21
	goto L9
L8:
	;
	v26 = int32(4)
	goto L9
L9:
	;
	v27 = v11 + v26
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v20 != v28 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	goto L12
L12:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+1)))
	if base.Ui32(v33) <= base.Ui32(v32) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int32(0)
L14:
	;
	goto L15
L15:
	;
	v37 = int32(2)
	v38 = v19 + v37
	v40 = v27 + v37
	v42 = int32(base.Ui32(v32) >> (uint(int32(3)) % 32))
	if base.Ui32(int32(4)) <= base.Ui32(v42) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	if v104 != 0 {
		goto L34
	} else {
		goto L35
	}
L17:
	;
	v104 = int32(0)
	goto L16
L18:
	;
	v78 = v73
	v79 = v74
	v80 = v75
	goto L28
L19:
	;
	if (v38|v40)&int32(3) != 0 {
		v73 = v38
		v74 = v40
		v75 = v42
		goto L18
	} else {
		goto L22
	}
L20:
	;
	v66 = v38
	v67 = v40
	v68 = v42
	goto L21
L21:
	;
	if v68 == int32(0) {
		goto L17
	} else {
		goto L27
	}
L22:
	;
	v50 = v38
	v51 = v40
	v52 = v42
	goto L23
L23:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	if v55 != v56 {
		v73 = v50
		v74 = v51
		v75 = v52
		goto L18
	} else {
		goto L25
	}
L24:
	;
	v66 = v61
	v67 = v59
	v68 = v63
	goto L21
L25:
	;
	v58 = int32(4)
	v59 = v51 + v58
	v61 = v50 + v58
	v63 = v52 - v58
	if base.Ui32(int32(3)) < base.Ui32(v63) {
		v50 = v61
		v51 = v59
		v52 = v63
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v73 = v66
	v74 = v67
	v75 = v68
	goto L18
L28:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	if v83 == v84 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v104 = v83 - v84
	goto L16
L30:
	;
	v86 = int32(1)
	v91 = v80 - v86
	if v91 != 0 {
		v78 = v78 + v86
		v79 = v79 + v86
		v80 = v91
		goto L28
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	goto L29
L33:
	;
	goto L17
L34:
	;
	return int32(0)
L35:
	;
	goto L36
L36:
	;
	v108 = v32 & int32(7)
	if v108 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	return int32(1)
L38:
	;
	goto L39
L39:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+v38))))
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+v40))))
	v117 = v114 ^ v116
	if base.Ui32(int32(127)) < base.Ui32(v117) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	return int32(0)
L41:
	;
	goto L42
L42:
	;
	v122 = int32(1)
	if v108 == v122 {
		v169 = v122
		goto L43
	} else {
		goto L44
	}
L43:
	;
	return v169
L44:
	;
	if v117<<(uint(int32(1))%32)&int32(128) != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	return int32(0)
L46:
	;
	goto L47
L47:
	;
	if base.Ui32(v108) < base.Ui32(int32(3)) {
		v169 = v122
		goto L43
	} else {
		goto L48
	}
L48:
	;
	if v117<<(uint(int32(2))%32)&int32(128) != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	return int32(0)
L50:
	;
	goto L51
L51:
	;
	if v108 == int32(3) {
		v169 = v122
		goto L43
	} else {
		goto L52
	}
L52:
	;
	if v117<<(uint(int32(3))%32)&int32(128) != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	return int32(0)
L54:
	;
	goto L55
L55:
	;
	if base.Ui32(v108) < base.Ui32(int32(5)) {
		v169 = v122
		goto L43
	} else {
		goto L56
	}
L56:
	;
	if v117<<(uint(int32(4))%32)&int32(128) != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	return int32(0)
L58:
	;
	goto L59
L59:
	;
	if v108 == int32(5) {
		v169 = v122
		goto L43
	} else {
		goto L60
	}
L60:
	;
	if v117<<(uint(int32(5))%32)&int32(128) != 0 {
		v169 = int32(0)
		goto L43
	} else {
		goto L61
	}
L61:
	;
	if v108 != int32(7) {
		v169 = int32(1)
		goto L43
	} else {
		goto L62
	}
L62:
	;
	v169 = base.B2i32(v117&int32(2) == int32(0))
	goto L43
}
func F_network_supeq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
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
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v169 int32
	_ = v169
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v13 = int32(1)
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if v15&v13 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v18 = v13
	goto L6
L5:
	;
	v18 = int32(4)
	goto L6
L6:
	;
	v19 = v6 + v18
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	v21 = int32(1)
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v23&v21 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v26 = v21
	goto L9
L8:
	;
	v26 = int32(4)
	goto L9
L9:
	;
	v27 = v11 + v26
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v20 != v28 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	goto L12
L12:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+1)))
	if base.Ui32(v33) < base.Ui32(v32) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int32(0)
L14:
	;
	goto L15
L15:
	;
	v37 = int32(2)
	v38 = v19 + v37
	v40 = v27 + v37
	v42 = int32(base.Ui32(v32) >> (uint(int32(3)) % 32))
	if base.Ui32(int32(4)) <= base.Ui32(v42) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	if v104 != 0 {
		goto L34
	} else {
		goto L35
	}
L17:
	;
	v104 = int32(0)
	goto L16
L18:
	;
	v78 = v73
	v79 = v74
	v80 = v75
	goto L28
L19:
	;
	if (v38|v40)&int32(3) != 0 {
		v73 = v38
		v74 = v40
		v75 = v42
		goto L18
	} else {
		goto L22
	}
L20:
	;
	v66 = v38
	v67 = v40
	v68 = v42
	goto L21
L21:
	;
	if v68 == int32(0) {
		goto L17
	} else {
		goto L27
	}
L22:
	;
	v50 = v38
	v51 = v40
	v52 = v42
	goto L23
L23:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	if v55 != v56 {
		v73 = v50
		v74 = v51
		v75 = v52
		goto L18
	} else {
		goto L25
	}
L24:
	;
	v66 = v61
	v67 = v59
	v68 = v63
	goto L21
L25:
	;
	v58 = int32(4)
	v59 = v51 + v58
	v61 = v50 + v58
	v63 = v52 - v58
	if base.Ui32(int32(3)) < base.Ui32(v63) {
		v50 = v61
		v51 = v59
		v52 = v63
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v73 = v66
	v74 = v67
	v75 = v68
	goto L18
L28:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	if v83 == v84 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v104 = v83 - v84
	goto L16
L30:
	;
	v86 = int32(1)
	v91 = v80 - v86
	if v91 != 0 {
		v78 = v78 + v86
		v79 = v79 + v86
		v80 = v91
		goto L28
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	goto L29
L33:
	;
	goto L17
L34:
	;
	return int32(0)
L35:
	;
	goto L36
L36:
	;
	v108 = v32 & int32(7)
	if v108 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	return int32(1)
L38:
	;
	goto L39
L39:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+v38))))
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+v40))))
	v117 = v114 ^ v116
	if base.Ui32(int32(127)) < base.Ui32(v117) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	return int32(0)
L41:
	;
	goto L42
L42:
	;
	v122 = int32(1)
	if v108 == v122 {
		v169 = v122
		goto L43
	} else {
		goto L44
	}
L43:
	;
	return v169
L44:
	;
	if v117<<(uint(int32(1))%32)&int32(128) != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	return int32(0)
L46:
	;
	goto L47
L47:
	;
	if base.Ui32(v108) < base.Ui32(int32(3)) {
		v169 = v122
		goto L43
	} else {
		goto L48
	}
L48:
	;
	if v117<<(uint(int32(2))%32)&int32(128) != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	return int32(0)
L50:
	;
	goto L51
L51:
	;
	if v108 == int32(3) {
		v169 = v122
		goto L43
	} else {
		goto L52
	}
L52:
	;
	if v117<<(uint(int32(3))%32)&int32(128) != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	return int32(0)
L54:
	;
	goto L55
L55:
	;
	if base.Ui32(v108) < base.Ui32(int32(5)) {
		v169 = v122
		goto L43
	} else {
		goto L56
	}
L56:
	;
	if v117<<(uint(int32(4))%32)&int32(128) != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	return int32(0)
L58:
	;
	goto L59
L59:
	;
	if v108 == int32(5) {
		v169 = v122
		goto L43
	} else {
		goto L60
	}
L60:
	;
	if v117<<(uint(int32(5))%32)&int32(128) != 0 {
		v169 = int32(0)
		goto L43
	} else {
		goto L61
	}
L61:
	;
	if v108 != int32(7) {
		v169 = int32(1)
		goto L43
	} else {
		goto L62
	}
L62:
	;
	v169 = base.B2i32(v117&int32(2) == int32(0))
	goto L43
}
