package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_calc_hist_selectivity_scalar(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) float64 {
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 float64
	_ = v54
	var v55 float64
	_ = v55
	var v61 int32
	_ = v61
	var v64 float64
	_ = v64
	var v65 int32
	_ = v65
	var v69 float64
	_ = v69
	v14 = l3 - int32(1)
	v20 = int32(-1)
	v21 = v14
	goto L1
L1:
	;
	v30 = base.I32_div_s(v20+v21+int32(1), int32(2))
	v34 = F_range_cmp_bounds(m, l0, l2+v30<<(uint(int32(3))%32), l1)
	v37 = m.ExcPending
	if v37 != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	v49 = int32(0)
	if v49 < v44 {
		goto L12
	} else {
		goto L13
	}
L3:
	;
	return float64(0)
L4:
	;
	v38 = int32(0)
	v43 = base.B2i32(v34 < v38) | l4&base.B2i32(v34 == v38)
	if v43 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v44 = v30
	goto L7
L6:
	;
	v44 = v20
	goto L7
L7:
	;
	if v43 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v47 = v21
	goto L10
L9:
	;
	v47 = v30 - int32(1)
	goto L10
L10:
	;
	if v44 < v47 {
		v20 = v44
		v21 = v47
		goto L1
	} else {
		goto L11
	}
L11:
	;
	goto L2
L12:
	;
	v52 = v44
	goto L14
L13:
	;
	v52 = v49
	goto L14
L14:
	;
	v54 = base.F64_convert_i32_u(v14)
	v55 = base.F64_div(base.F64_convert_i32_u(v52), v54)
	if v44 < int32(0) {
		v69 = v55
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return v69
L16:
	;
	if v14 <= v44 {
		v69 = v55
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v61 = l2 + v44<<(uint(int32(3))%32)
	v64 = F_get_position(m, l0, l1, v61, v61+int32(8))
	v65 = m.ExcPending
	if v65 != 0 {
		goto L3
	} else {
		goto L18
	}
L18:
	;
	v69 = base.F64_add(v55, base.F64_div(v64, v54))
	goto L15
}
func F_calc_key_id(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
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
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int64
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v13 = F_pgp_load_digest(m, int32(2), v8+int32(28))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		if int32(0) <= v13 {
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
			switch v22 - int32(1) {
			case 0, 1, 2:
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
				v55 = v36 + v38 + int32(10)
			default:
				v55 = int32(6)
			case 15:
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
				v55 = v26 + (v28 + v30) + int32(12)
			case 16:
				v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
				v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
				v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
				v55 = v43 + (v45 + (v47 + v49)) + int32(14)
			}
			v56 = int32(153)
			*(*uint8)(unsafe.Add(mBase, uint32(v8)+25)) = uint8(v56)
			*(*uint8)(unsafe.Add(mBase, uint32(v8)+27)) = uint8(v55)
			v59 = int32(8)
			v61 = int32(base.Ui32(v55) >> (uint(v59) % 32))
			*(*uint8)(unsafe.Add(mBase, uint32(v8)+26)) = uint8(v61)
			v63 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
			v67 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
			m.T0[v67].(func(*base.Module, int32, int32, int32))(m, v63, v8+int32(25), int32(3))
			mBase = m.M
			v69 = m.ExcPending
			if v69 != 0 {
				return int32(0)
			} else {
				v70 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
				v72 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
				m.T0[v72].(func(*base.Module, int32, int32, int32))(m, v70, l0, int32(1))
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return int32(0)
				} else {
					v75 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
					v79 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
					m.T0[v79].(func(*base.Module, int32, int32, int32))(m, v75, l0+int32(1), int32(4))
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return int32(0)
					} else {
						v82 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
						v84 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
						m.T0[v84].(func(*base.Module, int32, int32, int32))(m, v82, l0+int32(5), int32(1))
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
							return int32(0)
						} else {
							v87 = int32(12)
							v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
							switch v88 - int32(1) {
							case 0, 1, 2:
								v106 = v87
								v108 = v59
								v109 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
								v111 = *(*int32)(unsafe.Add(mBase, uint32(l0+v108)))
								v112 = F_pgp_mpi_hash(m, v109, v111)
								mBase = m.M
								v113 = m.ExcPending
								if v113 != 0 {
									return int32(0)
								} else {
									v114 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
									v116 = *(*int32)(unsafe.Add(mBase, uint32(l0+v106)))
									v117 = F_pgp_mpi_hash(m, v114, v116)
									mBase = m.M
									v118 = m.ExcPending
									if v118 != 0 {
										return int32(0)
									} else {
										v122 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
										v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+16))
										m.T0[v123].(func(*base.Module, int32, int32))(m, v122, v8)
										mBase = m.M
										v125 = m.ExcPending
										if v125 != 0 {
											return int32(0)
										} else {
											v126 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
											v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+20))
											m.T0[v127].(func(*base.Module, int32))(m, v126)
											mBase = m.M
											v129 = m.ExcPending
											if v129 != 0 {
												return int32(0)
											} else {
												v130 = *(*int64)(unsafe.Add(mBase, uint32(v8)+12))
												*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v130
												v132 = int32(0)
												v135 = F___memset(m, v8, v132, int32(20))
												mBase = m.M
												v136 = v132
												m.G0 = v8 + int32(32)
												return v136
											}
										}
									}
								}
							default:
								v122 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
								v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+16))
								m.T0[v123].(func(*base.Module, int32, int32))(m, v122, v8)
								mBase = m.M
								v125 = m.ExcPending
								if v125 != 0 {
									return int32(0)
								} else {
									v126 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
									v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+20))
									m.T0[v127].(func(*base.Module, int32))(m, v126)
									mBase = m.M
									v129 = m.ExcPending
									if v129 != 0 {
										return int32(0)
									} else {
										v130 = *(*int64)(unsafe.Add(mBase, uint32(v8)+12))
										*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v130
										v132 = int32(0)
										v135 = F___memset(m, v8, v132, int32(20))
										mBase = m.M
										v136 = v132
										m.G0 = v8 + int32(32)
										return v136
									}
								}
							case 15:
								v98 = v87
								v99 = v88
								v100 = v59
								v101 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
								v103 = *(*int32)(unsafe.Add(mBase, uint32(l0+v100)))
								v104 = F_pgp_mpi_hash(m, v101, v103)
								mBase = m.M
								v105 = m.ExcPending
								if v105 != 0 {
									return int32(0)
								} else {
									v106 = v99
									v108 = v98
									v109 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
									v111 = *(*int32)(unsafe.Add(mBase, uint32(l0+v108)))
									v112 = F_pgp_mpi_hash(m, v109, v111)
									mBase = m.M
									v113 = m.ExcPending
									if v113 != 0 {
										return int32(0)
									} else {
										v114 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
										v116 = *(*int32)(unsafe.Add(mBase, uint32(l0+v106)))
										v117 = F_pgp_mpi_hash(m, v114, v116)
										mBase = m.M
										v118 = m.ExcPending
										if v118 != 0 {
											return int32(0)
										} else {
											v122 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
											v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+16))
											m.T0[v123].(func(*base.Module, int32, int32))(m, v122, v8)
											mBase = m.M
											v125 = m.ExcPending
											if v125 != 0 {
												return int32(0)
											} else {
												v126 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
												v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+20))
												m.T0[v127].(func(*base.Module, int32))(m, v126)
												mBase = m.M
												v129 = m.ExcPending
												if v129 != 0 {
													return int32(0)
												} else {
													v130 = *(*int64)(unsafe.Add(mBase, uint32(v8)+12))
													*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v130
													v132 = int32(0)
													v135 = F___memset(m, v8, v132, int32(20))
													mBase = m.M
													v136 = v132
													m.G0 = v8 + int32(32)
													return v136
												}
											}
										}
									}
								}
							case 16:
								v91 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
								v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v93 = F_pgp_mpi_hash(m, v91, v92)
								mBase = m.M
								v94 = m.ExcPending
								if v94 != 0 {
									return int32(0)
								} else {
									v98 = int32(16)
									v99 = int32(20)
									v100 = int32(12)
									v101 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
									v103 = *(*int32)(unsafe.Add(mBase, uint32(l0+v100)))
									v104 = F_pgp_mpi_hash(m, v101, v103)
									mBase = m.M
									v105 = m.ExcPending
									if v105 != 0 {
										return int32(0)
									} else {
										v106 = v99
										v108 = v98
										v109 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
										v111 = *(*int32)(unsafe.Add(mBase, uint32(l0+v108)))
										v112 = F_pgp_mpi_hash(m, v109, v111)
										mBase = m.M
										v113 = m.ExcPending
										if v113 != 0 {
											return int32(0)
										} else {
											v114 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
											v116 = *(*int32)(unsafe.Add(mBase, uint32(l0+v106)))
											v117 = F_pgp_mpi_hash(m, v114, v116)
											mBase = m.M
											v118 = m.ExcPending
											if v118 != 0 {
												return int32(0)
											} else {
												v122 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
												v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+16))
												m.T0[v123].(func(*base.Module, int32, int32))(m, v122, v8)
												mBase = m.M
												v125 = m.ExcPending
												if v125 != 0 {
													return int32(0)
												} else {
													v126 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
													v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+20))
													m.T0[v127].(func(*base.Module, int32))(m, v126)
													mBase = m.M
													v129 = m.ExcPending
													if v129 != 0 {
														return int32(0)
													} else {
														v130 = *(*int64)(unsafe.Add(mBase, uint32(v8)+12))
														*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v130
														v132 = int32(0)
														v135 = F___memset(m, v8, v132, int32(20))
														mBase = m.M
														v136 = v132
														m.G0 = v8 + int32(32)
														return v136
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
			v136 = v13
			m.G0 = v8 + int32(32)
			return v136
		}
	}
}
func F_case_index(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	v2 = int32(0)
	if base.Ui32(l0) < base.Ui32(int32(1416)) {
		v89 = l0
		v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89<<(uint(int32(1))%32))+uint32(_c_F_case_index[0]))))
		v96 = v94
	} else {
		if base.Ui32(l0) <= base.Ui32(int32(_a_F_case_index_0)) {
			if base.Ui32(l0) <= base.Ui32(int32(_a_F_case_index_1)) {
				if base.Ui32(l0-int32(_a_F_case_index_2)) <= base.Ui32(int32(95)) {
					v89 = l0 - int32(2840)
					v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89<<(uint(int32(1))%32))+uint32(_c_F_case_index[0]))))
					v96 = v94
				} else {
					if base.Ui32(l0) < base.Ui32(int32(_a_F_case_index_3)) {
						v96 = v2
					} else {
						if base.Ui32(l0) <= base.Ui32(int32(_a_F_case_index_4)) {
							v89 = l0 - int32(3512)
							v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89<<(uint(int32(1))%32))+uint32(_c_F_case_index[0]))))
							v96 = v94
						} else {
							if base.Ui32(l0) < base.Ui32(int32(_a_F_case_index_5)) {
								v96 = v2
							} else {
								v89 = l0 - int32(_a_F_case_index_6)
								v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89<<(uint(int32(1))%32))+uint32(_c_F_case_index[0]))))
								v96 = v94
							}
						}
					}
				}
			} else {
				if base.Ui32(l0) < base.Ui32(int32(_a_F_case_index_7)) {
					v96 = v2
				} else {
					if base.Ui32(l0) <= base.Ui32(int32(_a_F_case_index_8)) {
						if base.Ui32(l0) <= base.Ui32(int32(_a_F_case_index_9)) {
							v89 = l0 - int32(_a_F_case_index_10)
							v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89<<(uint(int32(1))%32))+uint32(_c_F_case_index[0]))))
							v96 = v94
						} else {
							if base.Ui32(l0) < base.Ui32(int32(_a_F_case_index_11)) {
								v96 = v2
							} else {
								v89 = l0 - int32(_a_F_case_index_12)
								v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89<<(uint(int32(1))%32))+uint32(_c_F_case_index[0]))))
								v96 = v94
							}
						}
					} else {
						if base.Ui32(l0) < base.Ui32(int32(_a_F_case_index_13)) {
							v96 = v2
						} else {
							if base.Ui32(l0) <= base.Ui32(int32(_a_F_case_index_14)) {
								v89 = l0 - int32(_a_F_case_index_15)
								v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89<<(uint(int32(1))%32))+uint32(_c_F_case_index[0]))))
								v96 = v94
							} else {
								if base.Ui32(l0) < base.Ui32(int32(_a_F_case_index_16)) {
									v96 = v2
								} else {
									v89 = l0 - int32(_a_F_case_index_17)
									v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89<<(uint(int32(1))%32))+uint32(_c_F_case_index[0]))))
									v96 = v94
								}
							}
						}
					}
				}
			}
		} else {
			if base.Ui32(l0) < base.Ui32(int32(_a_F_case_index_18)) {
				v96 = v2
			} else {
				if base.Ui32(l0) <= base.Ui32(int32(_a_F_case_index_19)) {
					if base.Ui32(l0) <= base.Ui32(int32(_a_F_case_index_20)) {
						if base.Ui32(l0) <= base.Ui32(int32(_a_F_case_index_21)) {
							v89 = l0 - int32(_a_F_case_index_22)
							v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89<<(uint(int32(1))%32))+uint32(_c_F_case_index[0]))))
							v96 = v94
						} else {
							if base.Ui32(l0) < base.Ui32(int32(_a_F_case_index_23)) {
								v96 = v2
							} else {
								v89 = l0 - int32(_a_F_case_index_24)
								v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89<<(uint(int32(1))%32))+uint32(_c_F_case_index[0]))))
								v96 = v94
							}
						}
					} else {
						if base.Ui32(l0) < base.Ui32(int32(_a_F_case_index_25)) {
							v96 = v2
						} else {
							if base.Ui32(l0) <= base.Ui32(int32(_a_F_case_index_26)) {
								v89 = l0 - int32(_a_F_case_index_27)
								v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89<<(uint(int32(1))%32))+uint32(_c_F_case_index[0]))))
								v96 = v94
							} else {
								if base.Ui32(l0) < base.Ui32(int32(_a_F_case_index_28)) {
									v96 = v2
								} else {
									v89 = l0 - int32(_a_F_case_index_29)
									v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89<<(uint(int32(1))%32))+uint32(_c_F_case_index[0]))))
									v96 = v94
								}
							}
						}
					}
				} else {
					if base.Ui32(l0) < base.Ui32(int32(_a_F_case_index_30)) {
						v96 = v2
					} else {
						if base.Ui32(l0) <= base.Ui32(int32(_a_F_case_index_31)) {
							if base.Ui32(l0) <= base.Ui32(int32(_a_F_case_index_32)) {
								v89 = l0 - int32(_a_F_case_index_33)
								v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89<<(uint(int32(1))%32))+uint32(_c_F_case_index[0]))))
								v96 = v94
							} else {
								if base.Ui32(l0) < base.Ui32(int32(_a_F_case_index_34)) {
									v96 = v2
								} else {
									v89 = l0 - int32(_a_F_case_index_35)
									v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89<<(uint(int32(1))%32))+uint32(_c_F_case_index[0]))))
									v96 = v94
								}
							}
						} else {
							if base.Ui32(int32(67)) < base.Ui32(l0-int32(_a_F_case_index_36)) {
								v96 = v2
							} else {
								v89 = l0 - int32(_a_F_case_index_37)
								v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89<<(uint(int32(1))%32))+uint32(_c_F_case_index[0]))))
								v96 = v94
							}
						}
					}
				}
			}
		}
	}
	return v96
}
func F_catalan_ISO_8859_1_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v107 int32
	_ = v107
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v164 int32
	_ = v164
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v221 int32
	_ = v221
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v5))) = v6
	*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v6
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v18 < v9 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v9
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v239
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v239
	v243 = v239 - int32(1)
	if v243 <= v9 {
		goto L65
	} else {
		goto L66
	}
L2:
	;
	if v58 < int32(0) {
		goto L1
	} else {
		goto L17
	}
L3:
	;
	v20 = v9
	goto L5
L4:
	;
	v20 = v18
	goto L5
L5:
	;
	v27 = v9
	goto L7
L6:
	;
	v58 = v38
	goto L2
L7:
	;
	if v27 == v20 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v58 = int32(-1)
	goto L2
L10:
	;
	goto L11
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31+v27))))
	if int32(252) < v33 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v50 = v27 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v50
	v27 = v50
	goto L7
L13:
	;
	v35 = v33 - int32(97)
	if v35 < int32(0) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v38 = int32(1)
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v35)>>(uint(int32(3))%32)))+uint32(_c_F_catalan_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v42)>>(uint(v35&int32(7))%32))&v38 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	goto L12
L17:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v62 = v61 + v58
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v62
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v73 < v62 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if v116 < int32(0) {
		goto L1
	} else {
		goto L32
	}
L19:
	;
	v75 = v62
	goto L21
L20:
	;
	v75 = v73
	goto L21
L21:
	;
	v82 = v62
	goto L23
L22:
	;
	v116 = int32(1)
	goto L18
L23:
	;
	if v82 == v75 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v116 = int32(-1)
	goto L18
L26:
	;
	goto L27
L27:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88+v82))))
	if int32(252) < v90 {
		goto L22
	} else {
		goto L28
	}
L28:
	;
	v92 = v90 - int32(97)
	if v92 < int32(0) {
		goto L22
	} else {
		goto L29
	}
L29:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v92)>>(uint(int32(3))%32)))+uint32(_c_F_catalan_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v98)>>(uint(v92&int32(7))%32))&int32(1) == int32(0) {
		goto L22
	} else {
		goto L30
	}
L30:
	;
	v107 = v82 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v107
	v82 = v107
	goto L23
L32:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v120 = v119 + v116
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v120
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v122)+4)) = v120
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v132 < v131 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	if v172 < int32(0) {
		goto L1
	} else {
		goto L48
	}
L34:
	;
	v134 = v131
	goto L36
L35:
	;
	v134 = v132
	goto L36
L36:
	;
	v141 = v131
	goto L38
L37:
	;
	v172 = v152
	goto L33
L38:
	;
	if v141 == v134 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v172 = int32(-1)
	goto L33
L41:
	;
	goto L42
L42:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145+v141))))
	if int32(252) < v147 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v164 = v141 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v164
	v141 = v164
	goto L38
L44:
	;
	v149 = v147 - int32(97)
	if v149 < int32(0) {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v152 = int32(1)
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v149)>>(uint(int32(3))%32)))+uint32(_c_F_catalan_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v156)>>(uint(v149&int32(7))%32))&v152 != 0 {
		goto L37
	} else {
		goto L46
	}
L46:
	;
	goto L43
L48:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v176 = v175 + v172
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v176
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v187 < v176 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	if v230 < int32(0) {
		goto L1
	} else {
		goto L63
	}
L50:
	;
	v189 = v176
	goto L52
L51:
	;
	v189 = v187
	goto L52
L52:
	;
	v196 = v176
	goto L54
L53:
	;
	v230 = int32(1)
	goto L49
L54:
	;
	if v196 == v189 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v230 = int32(-1)
	goto L49
L57:
	;
	goto L58
L58:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202+v196))))
	if int32(252) < v204 {
		goto L53
	} else {
		goto L59
	}
L59:
	;
	v206 = v204 - int32(97)
	if v206 < int32(0) {
		goto L53
	} else {
		goto L60
	}
L60:
	;
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v206)>>(uint(int32(3))%32)))+uint32(_c_F_catalan_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v212)>>(uint(v206&int32(7))%32))&int32(1) == int32(0) {
		goto L53
	} else {
		goto L61
	}
L61:
	;
	v221 = v196 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v221
	v196 = v221
	goto L54
L63:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v233))) = v234 + v230
	goto L1
L64:
	;
	return v452
L65:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v276
	v281 = F_find_among_b(m, l0, int32(_a_F_catalan_ISO_8859_1_stem_0), int32(200))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L69
	} else {
		goto L77
	}
L66:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245+v243))))
	if v247&int32(224) != int32(96) {
		goto L65
	} else {
		goto L67
	}
L67:
	;
	if int32(1)<<(uint(v247)%32)&int32(_a_F_catalan_ISO_8859_1_stem_1) == int32(0) {
		goto L65
	} else {
		goto L68
	}
L68:
	;
	v260 = F_find_among_b(m, l0, int32(_a_F_catalan_ISO_8859_1_stem_2), int32(39))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	return int32(0)
L70:
	;
	if v260 == int32(0) {
		goto L65
	} else {
		goto L71
	}
L71:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v266
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v268)+4))
	if v266 < v269 {
		goto L65
	} else {
		goto L72
	}
L72:
	;
	v271 = F_slice_del(m, l0)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L69
	} else {
		goto L73
	}
L73:
	;
	if v271 < int32(0) {
		v452 = v271
		goto L64
	} else {
		goto L74
	}
L74:
	;
	goto L65
L75:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v360
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v360
	v365 = F_find_among_b(m, l0, int32(_a_F_catalan_ISO_8859_1_stem_3), int32(22))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L69
	} else {
		goto L110
	}
L76:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v331
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v331
	v336 = F_find_among_b(m, l0, int32(_a_F_catalan_ISO_8859_1_stem_4), int32(283))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L69
	} else {
		goto L99
	}
L77:
	;
	if v281 == int32(0) {
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v285
	switch v281 - int32(1) {
	case 0:
		goto L83
	case 1:
		goto L82
	case 2:
		goto L81
	case 3:
		goto L80
	case 4:
		goto L79
	default:
		goto L75
	}
L79:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v321)+4))
	if v285 < v322 {
		goto L76
	} else {
		goto L96
	}
L80:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v312)))
	if v285 < v313 {
		goto L76
	} else {
		goto L93
	}
L81:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v303)))
	if v285 < v304 {
		goto L76
	} else {
		goto L90
	}
L82:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v296)))
	if v285 < v297 {
		goto L76
	} else {
		goto L87
	}
L83:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v289)+4))
	if v285 < v290 {
		goto L76
	} else {
		goto L84
	}
L84:
	;
	v292 = F_slice_del(m, l0)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L69
	} else {
		goto L85
	}
L85:
	;
	if int32(0) <= v292 {
		goto L75
	} else {
		goto L86
	}
L86:
	;
	v452 = v292
	goto L64
L87:
	;
	v299 = F_slice_del(m, l0)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L69
	} else {
		goto L88
	}
L88:
	;
	if int32(0) <= v299 {
		goto L75
	} else {
		goto L89
	}
L89:
	;
	v452 = v299
	goto L64
L90:
	;
	v308 = F_slice_from_s(m, l0, int32(3), int32(_a_F_catalan_ISO_8859_1_stem_5))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L69
	} else {
		goto L91
	}
L91:
	;
	if int32(0) <= v308 {
		goto L75
	} else {
		goto L92
	}
L92:
	;
	v452 = v308
	goto L64
L93:
	;
	v317 = F_slice_from_s(m, l0, int32(2), int32(_a_F_catalan_ISO_8859_1_stem_6))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L69
	} else {
		goto L94
	}
L94:
	;
	if int32(0) <= v317 {
		goto L75
	} else {
		goto L95
	}
L95:
	;
	v452 = v317
	goto L64
L96:
	;
	v326 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_ISO_8859_1_stem_7))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L69
	} else {
		goto L97
	}
L97:
	;
	if int32(0) <= v326 {
		goto L75
	} else {
		goto L98
	}
L98:
	;
	v452 = v326
	goto L64
L99:
	;
	if v336 == int32(0) {
		goto L75
	} else {
		goto L100
	}
L100:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v340
	switch v336 - int32(1) {
	case 0:
		goto L102
	case 1:
		goto L101
	default:
		goto L75
	}
L101:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v351)))
	if v340 < v352 {
		goto L75
	} else {
		goto L106
	}
L102:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v344)+4))
	if v340 < v345 {
		goto L75
	} else {
		goto L103
	}
L103:
	;
	v347 = F_slice_del(m, l0)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L69
	} else {
		goto L104
	}
L104:
	;
	if int32(0) <= v347 {
		goto L75
	} else {
		goto L105
	}
L105:
	;
	v452 = v347
	goto L64
L106:
	;
	v354 = F_slice_del(m, l0)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L69
	} else {
		goto L107
	}
L107:
	;
	if v354 < int32(0) {
		v452 = v354
		goto L64
	} else {
		goto L108
	}
L108:
	;
	goto L75
L109:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v391
	v394 = v391
	goto L120
L110:
	;
	if v365 == int32(0) {
		goto L109
	} else {
		goto L111
	}
L111:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v369
	switch v365 - int32(1) {
	case 0:
		goto L113
	case 1:
		goto L112
	default:
		goto L109
	}
L112:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v380)+4))
	if v369 < v381 {
		goto L109
	} else {
		goto L117
	}
L113:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v373)+4))
	if v369 < v374 {
		goto L109
	} else {
		goto L114
	}
L114:
	;
	v376 = F_slice_del(m, l0)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L69
	} else {
		goto L115
	}
L115:
	;
	if int32(0) <= v376 {
		goto L109
	} else {
		goto L116
	}
L116:
	;
	v452 = v376
	goto L64
L117:
	;
	v385 = F_slice_from_s(m, l0, int32(2), int32(_a_F_catalan_ISO_8859_1_stem_8))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L69
	} else {
		goto L118
	}
L118:
	;
	if v385 < int32(0) {
		v452 = v385
		goto L64
	} else {
		goto L119
	}
L119:
	;
	goto L109
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v394
	v400 = F_find_among(m, l0, int32(_a_F_catalan_ISO_8859_1_stem_9), int32(13))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L69
	} else {
		goto L123
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v391
	v452 = int32(1)
	goto L64
L122:
	;
	goto L121
L123:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v402
	switch v400 - int32(1) {
	case 0:
		goto L131
	case 1:
		goto L130
	case 2:
		goto L129
	case 3:
		goto L128
	case 4:
		goto L127
	case 5:
		goto L126
	case 6:
		goto L125
	default:
		goto L124
	}
L124:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v394 = v448
	goto L120
L125:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v442 <= v402 {
		goto L122
	} else {
		goto L144
	}
L126:
	;
	v438 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_ISO_8859_1_stem_10))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L69
	} else {
		goto L142
	}
L127:
	;
	v432 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_ISO_8859_1_stem_11))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L69
	} else {
		goto L140
	}
L128:
	;
	v426 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_ISO_8859_1_stem_12))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L69
	} else {
		goto L138
	}
L129:
	;
	v420 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_ISO_8859_1_stem_13))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L69
	} else {
		goto L136
	}
L130:
	;
	v414 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_ISO_8859_1_stem_14))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L69
	} else {
		goto L134
	}
L131:
	;
	v408 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_ISO_8859_1_stem_15))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L69
	} else {
		goto L132
	}
L132:
	;
	if int32(0) <= v408 {
		goto L124
	} else {
		goto L133
	}
L133:
	;
	v452 = v408
	goto L64
L134:
	;
	if int32(0) <= v414 {
		goto L124
	} else {
		goto L135
	}
L135:
	;
	v452 = v414
	goto L64
L136:
	;
	if int32(0) <= v420 {
		goto L124
	} else {
		goto L137
	}
L137:
	;
	v452 = v420
	goto L64
L138:
	;
	if int32(0) <= v426 {
		goto L124
	} else {
		goto L139
	}
L139:
	;
	v452 = v426
	goto L64
L140:
	;
	if int32(0) <= v432 {
		goto L124
	} else {
		goto L141
	}
L141:
	;
	v452 = v432
	goto L64
L142:
	;
	if int32(0) <= v438 {
		goto L124
	} else {
		goto L143
	}
L143:
	;
	v452 = v438
	goto L64
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v402 + int32(1)
	goto L124
}
func F_catalan_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v83 int32
	_ = v83
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v205 int32
	_ = v205
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v239 int32
	_ = v239
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v279 int32
	_ = v279
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v330 int32
	_ = v330
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v356 int32
	_ = v356
	var v363 int32
	_ = v363
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v401 int32
	_ = v401
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v452 int32
	_ = v452
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v478 int32
	_ = v478
	var v486 int32
	_ = v486
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v627 int32
	_ = v627
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v754 int32
	_ = v754
	var v757 int32
	_ = v757
	var v761 int32
	_ = v761
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v783 int32
	_ = v783
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v795 int32
	_ = v795
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v7
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v7
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v32 = v10
	goto L4
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v10
	v506 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v506
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v506
	v510 = v506 - int32(1)
	if v510 <= v10 {
		goto L105
	} else {
		goto L106
	}
L2:
	;
	if v127 < int32(0) {
		goto L1
	} else {
		goto L27
	}
L3:
	;
	v127 = v99
	goto L2
L4:
	;
	if v23 <= v32 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v127 = int32(-1)
	goto L2
L7:
	;
	goto L8
L8:
	;
	v39 = int32(1)
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32+v24))))
	if base.Ui32(v41) < base.Ui32(int32(192)) {
		v98 = v41
		v99 = v39
		goto L9
	} else {
		goto L10
	}
L9:
	;
	if int32(252) < v98 {
		goto L22
	} else {
		goto L23
	}
L10:
	;
	v45 = v32 + int32(1)
	if v45 == v23 {
		v98 = v41
		v99 = v39
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45+v24))))
	v50 = v48 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v41) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54+v24))))
	v66 = v64 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v41) {
		goto L18
	} else {
		goto L19
	}
L13:
	;
	v54 = v32 + int32(2)
	if v54 != v23 {
		goto L12
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v98 = v41<<(uint(int32(6))%32)&int32(1984) | v50
	v99 = int32(2)
	goto L9
L16:
	;
	goto L15
L17:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v70))))
	v98 = v83&int32(63) | (v41<<(uint(int32(18))%32)&int32(_a_F_catalan_UTF_8_stem_0) | v50<<(uint(int32(12))%32) | v66<<(uint(int32(6))%32))
	v99 = int32(4)
	goto L9
L18:
	;
	v70 = v32 + int32(3)
	if v70 != v23 {
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v98 = v41<<(uint(int32(12))%32)&int32(_a_F_catalan_UTF_8_stem_1) | v50<<(uint(int32(6))%32) | v66
	v99 = int32(3)
	goto L9
L21:
	;
	goto L20
L22:
	;
	v116 = v99 + v32
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v116
	v32 = v116
	goto L4
L23:
	;
	v103 = v98 - int32(97)
	if v103 < int32(0) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v103)>>(uint(int32(3))%32)))+uint32(_c_F_catalan_UTF_8_stem[0]))))
	if int32(base.Ui32(v109)>>(uint(v103&int32(7))%32))&int32(1) != 0 {
		goto L3
	} else {
		goto L25
	}
L25:
	;
	goto L22
L27:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v131 = v130 + v127
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v131
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v154 = v131
	goto L30
L28:
	;
	if v250 < int32(0) {
		goto L1
	} else {
		goto L52
	}
L29:
	;
	v250 = v221
	goto L28
L30:
	;
	if v145 <= v154 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v250 = int32(-1)
	goto L28
L33:
	;
	goto L34
L34:
	;
	v161 = int32(1)
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154+v146))))
	if base.Ui32(v163) < base.Ui32(int32(192)) {
		v220 = v163
		v221 = v161
		goto L35
	} else {
		goto L36
	}
L35:
	;
	if int32(252) < v220 {
		goto L29
	} else {
		goto L48
	}
L36:
	;
	v167 = v154 + int32(1)
	if v167 == v145 {
		v220 = v163
		v221 = v161
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167+v146))))
	v172 = v170 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v163) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176+v146))))
	v188 = v186 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v163) {
		goto L44
	} else {
		goto L45
	}
L39:
	;
	v176 = v154 + int32(2)
	if v176 != v145 {
		goto L38
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v220 = v163<<(uint(int32(6))%32)&int32(1984) | v172
	v221 = int32(2)
	goto L35
L42:
	;
	goto L41
L43:
	;
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146+v192))))
	v220 = v205&int32(63) | (v163<<(uint(int32(18))%32)&int32(_a_F_catalan_UTF_8_stem_0) | v172<<(uint(int32(12))%32) | v188<<(uint(int32(6))%32))
	v221 = int32(4)
	goto L35
L44:
	;
	v192 = v154 + int32(3)
	if v192 != v145 {
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v220 = v163<<(uint(int32(12))%32)&int32(_a_F_catalan_UTF_8_stem_1) | v172<<(uint(int32(6))%32) | v188
	v221 = int32(3)
	goto L35
L47:
	;
	goto L46
L48:
	;
	v225 = v220 - int32(97)
	if v225 < int32(0) {
		goto L29
	} else {
		goto L49
	}
L49:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v225)>>(uint(int32(3))%32)))+uint32(_c_F_catalan_UTF_8_stem[0]))))
	if int32(base.Ui32(v231)>>(uint(v225&int32(7))%32))&int32(1) == int32(0) {
		goto L29
	} else {
		goto L50
	}
L50:
	;
	v239 = v221 + v154
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v239
	v154 = v239
	goto L30
L52:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v254 = v253 + v250
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v254
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v256)+4)) = v254
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v279 = v269
	goto L55
L53:
	;
	if v374 < int32(0) {
		goto L1
	} else {
		goto L78
	}
L54:
	;
	v374 = v346
	goto L53
L55:
	;
	if v270 <= v279 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v374 = int32(-1)
	goto L53
L58:
	;
	goto L59
L59:
	;
	v286 = int32(1)
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279+v271))))
	if base.Ui32(v288) < base.Ui32(int32(192)) {
		v345 = v288
		v346 = v286
		goto L60
	} else {
		goto L61
	}
L60:
	;
	if int32(252) < v345 {
		goto L73
	} else {
		goto L74
	}
L61:
	;
	v292 = v279 + int32(1)
	if v292 == v270 {
		v345 = v288
		v346 = v286
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292+v271))))
	v297 = v295 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v288) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301+v271))))
	v313 = v311 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v288) {
		goto L69
	} else {
		goto L70
	}
L64:
	;
	v301 = v279 + int32(2)
	if v301 != v270 {
		goto L63
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v345 = v288<<(uint(int32(6))%32)&int32(1984) | v297
	v346 = int32(2)
	goto L60
L67:
	;
	goto L66
L68:
	;
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271+v317))))
	v345 = v330&int32(63) | (v288<<(uint(int32(18))%32)&int32(_a_F_catalan_UTF_8_stem_0) | v297<<(uint(int32(12))%32) | v313<<(uint(int32(6))%32))
	v346 = int32(4)
	goto L60
L69:
	;
	v317 = v279 + int32(3)
	if v317 != v270 {
		goto L68
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v345 = v288<<(uint(int32(12))%32)&int32(_a_F_catalan_UTF_8_stem_1) | v297<<(uint(int32(6))%32) | v313
	v346 = int32(3)
	goto L60
L72:
	;
	goto L71
L73:
	;
	v363 = v346 + v279
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v363
	v279 = v363
	goto L55
L74:
	;
	v350 = v345 - int32(97)
	if v350 < int32(0) {
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v350)>>(uint(int32(3))%32)))+uint32(_c_F_catalan_UTF_8_stem[0]))))
	if int32(base.Ui32(v356)>>(uint(v350&int32(7))%32))&int32(1) != 0 {
		goto L54
	} else {
		goto L76
	}
L76:
	;
	goto L73
L78:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v378 = v377 + v374
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v378
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v401 = v378
	goto L81
L79:
	;
	if v497 < int32(0) {
		goto L1
	} else {
		goto L103
	}
L80:
	;
	v497 = v468
	goto L79
L81:
	;
	if v392 <= v401 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v497 = int32(-1)
	goto L79
L84:
	;
	goto L85
L85:
	;
	v408 = int32(1)
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v401+v393))))
	if base.Ui32(v410) < base.Ui32(int32(192)) {
		v467 = v410
		v468 = v408
		goto L86
	} else {
		goto L87
	}
L86:
	;
	if int32(252) < v467 {
		goto L80
	} else {
		goto L99
	}
L87:
	;
	v414 = v401 + int32(1)
	if v414 == v392 {
		v467 = v410
		v468 = v408
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v414+v393))))
	v419 = v417 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v410) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v423+v393))))
	v435 = v433 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v410) {
		goto L95
	} else {
		goto L96
	}
L90:
	;
	v423 = v401 + int32(2)
	if v423 != v392 {
		goto L89
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v467 = v410<<(uint(int32(6))%32)&int32(1984) | v419
	v468 = int32(2)
	goto L86
L93:
	;
	goto L92
L94:
	;
	v452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v393+v439))))
	v467 = v452&int32(63) | (v410<<(uint(int32(18))%32)&int32(_a_F_catalan_UTF_8_stem_0) | v419<<(uint(int32(12))%32) | v435<<(uint(int32(6))%32))
	v468 = int32(4)
	goto L86
L95:
	;
	v439 = v401 + int32(3)
	if v439 != v392 {
		goto L94
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v467 = v410<<(uint(int32(12))%32)&int32(_a_F_catalan_UTF_8_stem_1) | v419<<(uint(int32(6))%32) | v435
	v468 = int32(3)
	goto L86
L98:
	;
	goto L97
L99:
	;
	v472 = v467 - int32(97)
	if v472 < int32(0) {
		goto L80
	} else {
		goto L100
	}
L100:
	;
	v478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v472)>>(uint(int32(3))%32)))+uint32(_c_F_catalan_UTF_8_stem[0]))))
	if int32(base.Ui32(v478)>>(uint(v472&int32(7))%32))&int32(1) == int32(0) {
		goto L80
	} else {
		goto L101
	}
L101:
	;
	v486 = v468 + v401
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v486
	v401 = v486
	goto L81
L103:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v501 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v500))) = v501 + v497
	goto L1
L104:
	;
	return v795
L105:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v543
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v543
	v548 = F_find_among_b(m, l0, int32(_a_F_catalan_UTF_8_stem_2), int32(200))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L109
	} else {
		goto L117
	}
L106:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512+v510))))
	if v514&int32(224) != int32(96) {
		goto L105
	} else {
		goto L107
	}
L107:
	;
	if int32(1)<<(uint(v514)%32)&int32(_a_F_catalan_UTF_8_stem_3) == int32(0) {
		goto L105
	} else {
		goto L108
	}
L108:
	;
	v527 = F_find_among_b(m, l0, int32(_a_F_catalan_UTF_8_stem_4), int32(39))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	return int32(0)
L110:
	;
	if v527 == int32(0) {
		goto L105
	} else {
		goto L111
	}
L111:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v533
	v535 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v535)+4))
	if v533 < v536 {
		goto L105
	} else {
		goto L112
	}
L112:
	;
	v538 = F_slice_del(m, l0)
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L109
	} else {
		goto L113
	}
L113:
	;
	if v538 < int32(0) {
		v795 = v538
		goto L104
	} else {
		goto L114
	}
L114:
	;
	goto L105
L115:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v627
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v627
	v632 = F_find_among_b(m, l0, int32(_a_F_catalan_UTF_8_stem_5), int32(22))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L109
	} else {
		goto L150
	}
L116:
	;
	v598 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v598
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v598
	v603 = F_find_among_b(m, l0, int32(_a_F_catalan_UTF_8_stem_6), int32(283))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L109
	} else {
		goto L139
	}
L117:
	;
	if v548 == int32(0) {
		goto L116
	} else {
		goto L118
	}
L118:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v552
	switch v548 - int32(1) {
	case 0:
		goto L123
	case 1:
		goto L122
	case 2:
		goto L121
	case 3:
		goto L120
	case 4:
		goto L119
	default:
		goto L115
	}
L119:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v588)+4))
	if v552 < v589 {
		goto L116
	} else {
		goto L136
	}
L120:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v579)))
	if v552 < v580 {
		goto L116
	} else {
		goto L133
	}
L121:
	;
	v570 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v570)))
	if v552 < v571 {
		goto L116
	} else {
		goto L130
	}
L122:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v563)))
	if v552 < v564 {
		goto L116
	} else {
		goto L127
	}
L123:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v556)+4))
	if v552 < v557 {
		goto L116
	} else {
		goto L124
	}
L124:
	;
	v559 = F_slice_del(m, l0)
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L109
	} else {
		goto L125
	}
L125:
	;
	if int32(0) <= v559 {
		goto L115
	} else {
		goto L126
	}
L126:
	;
	v795 = v559
	goto L104
L127:
	;
	v566 = F_slice_del(m, l0)
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L109
	} else {
		goto L128
	}
L128:
	;
	if int32(0) <= v566 {
		goto L115
	} else {
		goto L129
	}
L129:
	;
	v795 = v566
	goto L104
L130:
	;
	v575 = F_slice_from_s(m, l0, int32(3), int32(_a_F_catalan_UTF_8_stem_7))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L109
	} else {
		goto L131
	}
L131:
	;
	if int32(0) <= v575 {
		goto L115
	} else {
		goto L132
	}
L132:
	;
	v795 = v575
	goto L104
L133:
	;
	v584 = F_slice_from_s(m, l0, int32(2), int32(_a_F_catalan_UTF_8_stem_8))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L109
	} else {
		goto L134
	}
L134:
	;
	if int32(0) <= v584 {
		goto L115
	} else {
		goto L135
	}
L135:
	;
	v795 = v584
	goto L104
L136:
	;
	v593 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_UTF_8_stem_9))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L109
	} else {
		goto L137
	}
L137:
	;
	if int32(0) <= v593 {
		goto L115
	} else {
		goto L138
	}
L138:
	;
	v795 = v593
	goto L104
L139:
	;
	if v603 == int32(0) {
		goto L115
	} else {
		goto L140
	}
L140:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v607
	switch v603 - int32(1) {
	case 0:
		goto L142
	case 1:
		goto L141
	default:
		goto L115
	}
L141:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v618)))
	if v607 < v619 {
		goto L115
	} else {
		goto L146
	}
L142:
	;
	v611 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v611)+4))
	if v607 < v612 {
		goto L115
	} else {
		goto L143
	}
L143:
	;
	v614 = F_slice_del(m, l0)
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L109
	} else {
		goto L144
	}
L144:
	;
	if int32(0) <= v614 {
		goto L115
	} else {
		goto L145
	}
L145:
	;
	v795 = v614
	goto L104
L146:
	;
	v621 = F_slice_del(m, l0)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L109
	} else {
		goto L147
	}
L147:
	;
	if v621 < int32(0) {
		v795 = v621
		goto L104
	} else {
		goto L148
	}
L148:
	;
	goto L115
L149:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v658
	v660 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v662 = v658
	v663 = v660
	goto L160
L150:
	;
	if v632 == int32(0) {
		goto L149
	} else {
		goto L151
	}
L151:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v636
	switch v632 - int32(1) {
	case 0:
		goto L153
	case 1:
		goto L152
	default:
		goto L149
	}
L152:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v647)+4))
	if v636 < v648 {
		goto L149
	} else {
		goto L157
	}
L153:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v640)+4))
	if v636 < v641 {
		goto L149
	} else {
		goto L154
	}
L154:
	;
	v643 = F_slice_del(m, l0)
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L109
	} else {
		goto L155
	}
L155:
	;
	if int32(0) <= v643 {
		goto L149
	} else {
		goto L156
	}
L156:
	;
	v795 = v643
	goto L104
L157:
	;
	v652 = F_slice_from_s(m, l0, int32(2), int32(_a_F_catalan_UTF_8_stem_10))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L109
	} else {
		goto L158
	}
L158:
	;
	if v652 < int32(0) {
		v795 = v652
		goto L104
	} else {
		goto L159
	}
L159:
	;
	goto L149
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v662
	v668 = v662 + int32(1)
	if v663 <= v668 {
		goto L166
	} else {
		goto L167
	}
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v658
	v795 = int32(1)
	goto L104
L162:
	;
	goto L161
L163:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v791 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v662 = v791
	v663 = v790
	goto L160
L164:
	;
	v731 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L192
L165:
	;
	v685 = F_find_among(m, l0, int32(_a_F_catalan_UTF_8_stem_11), int32(13))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L109
	} else {
		goto L170
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v662
	v728 = v662
	v729 = v663
	goto L164
L167:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v672 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v670+v668))))
	if v672&int32(224) != int32(160) {
		goto L166
	} else {
		goto L168
	}
L168:
	;
	if int32(1)<<(uint(v672)%32)&int32(344765187) != 0 {
		goto L165
	} else {
		goto L169
	}
L169:
	;
	goto L166
L170:
	;
	v687 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v687
	switch v685 - int32(1) {
	case 0:
		goto L177
	case 1:
		goto L176
	case 2:
		goto L175
	case 3:
		goto L174
	case 4:
		goto L173
	case 5:
		goto L172
	case 6:
		goto L171
	default:
		goto L163
	}
L171:
	;
	v727 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v728 = v687
	v729 = v727
	goto L164
L172:
	;
	v723 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_UTF_8_stem_12))
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L109
	} else {
		goto L188
	}
L173:
	;
	v717 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_UTF_8_stem_13))
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L109
	} else {
		goto L186
	}
L174:
	;
	v711 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_UTF_8_stem_14))
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L109
	} else {
		goto L184
	}
L175:
	;
	v705 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_UTF_8_stem_15))
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L109
	} else {
		goto L182
	}
L176:
	;
	v699 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_UTF_8_stem_16))
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L109
	} else {
		goto L180
	}
L177:
	;
	v693 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_UTF_8_stem_17))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L109
	} else {
		goto L178
	}
L178:
	;
	if int32(0) <= v693 {
		goto L163
	} else {
		goto L179
	}
L179:
	;
	v795 = v693
	goto L104
L180:
	;
	if int32(0) <= v699 {
		goto L163
	} else {
		goto L181
	}
L181:
	;
	v795 = v699
	goto L104
L182:
	;
	if int32(0) <= v705 {
		goto L163
	} else {
		goto L183
	}
L183:
	;
	v795 = v705
	goto L104
L184:
	;
	if int32(0) <= v711 {
		goto L163
	} else {
		goto L185
	}
L185:
	;
	v795 = v711
	goto L104
L186:
	;
	if int32(0) <= v717 {
		goto L163
	} else {
		goto L187
	}
L187:
	;
	v795 = v717
	goto L104
L188:
	;
	if int32(0) <= v723 {
		goto L163
	} else {
		goto L189
	}
L189:
	;
	v795 = v723
	goto L104
L190:
	;
	if v783 < int32(0) {
		goto L162
	} else {
		goto L210
	}
L192:
	;
	goto L193
L193:
	;
	goto L194
L194:
	;
	v738 = v728
	v740 = int32(1)
	goto L197
L196:
	;
	v783 = v768
	goto L190
L197:
	;
	if v729 <= v738 {
		goto L199
	} else {
		goto L200
	}
L198:
	;
	goto L196
L199:
	;
	v783 = int32(-1)
	goto L190
L200:
	;
	goto L201
L201:
	;
	v745 = v738 + int32(1)
	v747 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v731+v738))))
	if base.Ui32(v747) < base.Ui32(int32(192)) {
		v768 = v745
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v769 = int32(1)
	if v769 < v740 {
		v738 = v768
		v740 = v740 - v769
		goto L197
	} else {
		goto L209
	}
L203:
	;
	if v729 <= v745 {
		v768 = v745
		goto L202
	} else {
		goto L204
	}
L204:
	;
	v754 = v745
	goto L205
L205:
	;
	v757 = int32(*(*int8)(unsafe.Add(mBase, uint32(v731+v754))))
	if int32(-65) < v757 {
		v768 = v754
		goto L202
	} else {
		goto L207
	}
L206:
	;
	v768 = v729
	goto L202
L207:
	;
	v761 = v754 + int32(1)
	if v761 != v729 {
		v754 = v761
		goto L205
	} else {
		goto L208
	}
L208:
	;
	goto L206
L209:
	;
	goto L198
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v783
	goto L163
}
