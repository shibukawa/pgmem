package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_PredicateLockPage(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockPage[0]))
	if v11 == int32(0) {
		m.G0 = v8 + int32(16)
		return
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		if v14 != 0 {
			m.G0 = v8 + int32(16)
			return
		} else {
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+108)))
			if v15&int32(128) != 0 {
				F_ReleasePredicateLocks(m, int32(0), int32(1))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					m.G0 = v8 + int32(16)
					return
				}
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				if base.Ui32(v22) < base.Ui32(int32(_a_F_PredicateLockPage_0)) {
					m.G0 = v8 + int32(16)
					return
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+118)))
					if v26 == int32(116) {
						m.G0 = v8 + int32(16)
						return
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v22
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v29
						F_PredicateLockAcquire(m, v8)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return
						} else {
							m.G0 = v8 + int32(16)
							return
						}
					}
				}
			}
		}
	}
}
func F_PredicateLockPageSplit(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
	var v42 int64
	_ = v42
	var v44 int64
	_ = v44
	var v46 int64
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v63 int64
	_ = v63
	var v65 int64
	_ = v65
	var v67 int64
	_ = v67
	var v69 int64
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	v6 = m.G0
	v8 = v6 - int32(96)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockPageSplit[0]))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	if v12 == int32(0) {
		m.G0 = v8 + int32(96)
		return
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		if base.Ui32(v15) < base.Ui32(int32(_a_F_PredicateLockPageSplit_0)) {
			m.G0 = v8 + int32(96)
			return
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+118)))
			if v19 == int32(116) {
				m.G0 = v8 + int32(96)
				return
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v23 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v8)+92)) = v23
				*(*int32)(unsafe.Add(mBase, uint32(v8)+88)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v8)+84)) = v15
				*(*int32)(unsafe.Add(mBase, uint32(v8)+80)) = v22
				*(*int32)(unsafe.Add(mBase, uint32(v8)+76)) = v23
				*(*int32)(unsafe.Add(mBase, uint32(v8)+72)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v8)+68)) = v15
				*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = v22
				v34 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockPageSplit[1]))
				v38 = F_LWLockAcquire(m, v34+int32(3840), v23)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					v40 = *(*int64)(unsafe.Add(mBase, uint32(v8)+88))
					*(*int64)(unsafe.Add(mBase, uint32(v8)+56)) = v40
					v42 = *(*int64)(unsafe.Add(mBase, uint32(v8)+80))
					*(*int64)(unsafe.Add(mBase, uint32(v8)+48)) = v42
					v44 = *(*int64)(unsafe.Add(mBase, uint32(v8)+64))
					*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = v44
					v46 = *(*int64)(unsafe.Add(mBase, uint32(v8)+72))
					*(*int64)(unsafe.Add(mBase, uint32(v8)+40)) = v46
					v53 = F_TransferPredicateLocksToNewTarget(m, v8+int32(48), v8+int32(32), int32(0))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return
					} else {
						if v53 == int32(0) {
							if l1 != int32(-1) {
								*(*int64)(unsafe.Add(mBase, uint32(v8)+72)) = int64(4294967295)
								*(*int32)(unsafe.Add(mBase, uint32(v8)+68)) = v15
								*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = v22
							} else {
							}
							v63 = *(*int64)(unsafe.Add(mBase, uint32(v8)+88))
							*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v63
							v65 = *(*int64)(unsafe.Add(mBase, uint32(v8)+80))
							*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v65
							v67 = *(*int64)(unsafe.Add(mBase, uint32(v8)+64))
							*(*int64)(unsafe.Add(mBase, uint32(v8))) = v67
							v69 = *(*int64)(unsafe.Add(mBase, uint32(v8)+72))
							*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v69
							v74 = F_TransferPredicateLocksToNewTarget(m, v8+int32(16), v8, int32(1))
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return
							} else {
								v77 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockPageSplit[1]))
								F_LWLockRelease(m, v77+int32(3840))
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return
								} else {
									m.G0 = v8 + int32(96)
									return
								}
							}
						} else {
							v77 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockPageSplit[1]))
							F_LWLockRelease(m, v77+int32(3840))
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return
							} else {
								m.G0 = v8 + int32(96)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_PredicateLockShmemRequest(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
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
	var v18 int64
	_ = v18
	var v42 int64
	_ = v42
	var v47 int32
	_ = v47
	var v48 int64
	_ = v48
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v105 int32
	_ = v105
	var v106 int64
	_ = v106
	var v131 int64
	_ = v131
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v156 int32
	_ = v156
	var v166 int32
	_ = v166
	var v167 int64
	_ = v167
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v203 int32
	_ = v203
	v4 = m.G0
	v6 = v4 - int32(400)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockShmemRequest[0]))
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockShmemRequest[1]))
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockShmemRequest[2]))
	v14 = F_add_size(m, v11, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		v16 = F_mul_size(m, v9, v14)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			v18 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v6)+320)) = v18
			*(*int64)(unsafe.Add(mBase, uint32(v6)+312)) = v18
			*(*int64)(unsafe.Add(mBase, uint32(v6)+360)) = v18
			*(*int64)(unsafe.Add(mBase, uint32(v6)+352)) = int64(103079215120)
			*(*int64)(unsafe.Add(mBase, uint32(v6)+344)) = int64(16)
			*(*int32)(unsafe.Add(mBase, uint32(v6)+332)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v6)+328)) = int32(_a_F_PredicateLockShmemRequest_0)
			*(*int64)(unsafe.Add(mBase, uint32(v6)+368)) = v18
			*(*int64)(unsafe.Add(mBase, uint32(v6)+376)) = v18
			*(*int64)(unsafe.Add(mBase, uint32(v6)+384)) = v18
			*(*int32)(unsafe.Add(mBase, uint32(v6)+396)) = int32(_a_F_PredicateLockShmemRequest_1)
			*(*int32)(unsafe.Add(mBase, uint32(v6)+392)) = int32(_a_F_PredicateLockShmemRequest_2)
			v42 = base.I64_extend_i32_u(v16)
			*(*int64)(unsafe.Add(mBase, uint32(v6)+336)) = v42
			F_ShmemRequestHashWithOpts(m, v6+int32(312))
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return
			} else {
				v48 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v6)+232)) = v48
				*(*int64)(unsafe.Add(mBase, uint32(v6)+224)) = v48
				*(*int64)(unsafe.Add(mBase, uint32(v6)+276)) = v48
				*(*int32)(unsafe.Add(mBase, uint32(v6)+272)) = int32(1221)
				*(*int64)(unsafe.Add(mBase, uint32(v6)+264)) = int64(137438953480)
				*(*int64)(unsafe.Add(mBase, uint32(v6)+256)) = int64(16)
				*(*int64)(unsafe.Add(mBase, uint32(v6)+248)) = v42 << (uint(int64(1)) % 64)
				*(*int32)(unsafe.Add(mBase, uint32(v6)+244)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v6)+240)) = int32(_a_F_PredicateLockShmemRequest_3)
				*(*int64)(unsafe.Add(mBase, uint32(v6)+284)) = v48
				*(*int64)(unsafe.Add(mBase, uint32(v6)+292)) = v48
				*(*int32)(unsafe.Add(mBase, uint32(v6)+308)) = int32(_a_F_PredicateLockShmemRequest_4)
				*(*int64)(unsafe.Add(mBase, uint32(v6)+300)) = int64(35497904701440)
				F_ShmemRequestHashWithOpts(m, v6+int32(224))
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return
				} else {
					v81 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockShmemRequest[2]))
					v83 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockShmemRequest[1]))
					v86 = (v81 + v83) * int32(10)
					*(*int64)(unsafe.Add(mBase, _c_F_PredicateLockShmemRequest[3])) = base.I64_extend_i32_s(v86)
					*(*int32)(unsafe.Add(mBase, uint32(v6)+208)) = int32(_a_F_PredicateLockShmemRequest_5)
					v93 = F_mul_size(m, v86, int32(120))
					mBase = m.M
					v94 = m.ExcPending
					if v94 != 0 {
						return
					} else {
						v95 = F_add_size(m, int32(64), v93)
						mBase = m.M
						v96 = m.ExcPending
						if v96 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6)+220)) = int32(_a_F_PredicateLockShmemRequest_6)
							*(*int32)(unsafe.Add(mBase, uint32(v6)+216)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v6)+212)) = v95
							F_ShmemRequestStructWithOpts(m, v6+int32(208))
							mBase = m.M
							v105 = m.ExcPending
							if v105 != 0 {
								return
							} else {
								v106 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v6)+128)) = v106
								*(*int64)(unsafe.Add(mBase, uint32(v6)+120)) = v106
								*(*int32)(unsafe.Add(mBase, uint32(v6)+140)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v6)+136)) = int32(_a_F_PredicateLockShmemRequest_7)
								*(*int64)(unsafe.Add(mBase, uint32(v6)+152)) = v106
								*(*int64)(unsafe.Add(mBase, uint32(v6)+160)) = int64(34359738372)
								*(*int64)(unsafe.Add(mBase, uint32(v6)+168)) = v106
								*(*int64)(unsafe.Add(mBase, uint32(v6)+176)) = v106
								*(*int64)(unsafe.Add(mBase, uint32(v6)+184)) = v106
								*(*int64)(unsafe.Add(mBase, uint32(v6)+192)) = v106
								*(*int32)(unsafe.Add(mBase, uint32(v6)+204)) = int32(_a_F_PredicateLockShmemRequest_8)
								*(*int32)(unsafe.Add(mBase, uint32(v6)+200)) = int32(_a_F_PredicateLockShmemRequest_9)
								v131 = *(*int64)(unsafe.Add(mBase, _c_F_PredicateLockShmemRequest[3]))
								*(*int64)(unsafe.Add(mBase, uint32(v6)+144)) = v131
								F_ShmemRequestHashWithOpts(m, v6+int32(120))
								mBase = m.M
								v136 = m.ExcPending
								if v136 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v6)+104)) = int32(_a_F_PredicateLockShmemRequest_10)
									v140 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockShmemRequest[3]))
									v144 = F_mul_size(m, v140*int32(5), int32(24))
									mBase = m.M
									v145 = m.ExcPending
									if v145 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v6)+116)) = int32(_a_F_PredicateLockShmemRequest_11)
										*(*int32)(unsafe.Add(mBase, uint32(v6)+112)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v6)+108)) = v144 + int32(16)
										F_ShmemRequestStructWithOpts(m, v6+int32(104))
										mBase = m.M
										v156 = m.ExcPending
										if v156 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v6)+100)) = int32(_a_F_PredicateLockShmemRequest_12)
											*(*int64)(unsafe.Add(mBase, uint32(v6)+92)) = int64(8)
											*(*int32)(unsafe.Add(mBase, uint32(v6)+88)) = int32(_a_F_PredicateLockShmemRequest_13)
											F_ShmemRequestStructWithOpts(m, v6+int32(88))
											mBase = m.M
											v166 = m.ExcPending
											if v166 != 0 {
												return
											} else {
												v167 = int64(0)
												*(*int64)(unsafe.Add(mBase, uint32(v6)+32)) = v167
												*(*int64)(unsafe.Add(mBase, uint32(v6)+24)) = v167
												*(*int32)(unsafe.Add(mBase, uint32(v6)+44)) = int32(_a_F_PredicateLockShmemRequest_14)
												*(*int32)(unsafe.Add(mBase, uint32(v6)+40)) = int32(_a_F_PredicateLockShmemRequest_15)
												*(*int64)(unsafe.Add(mBase, uint32(v6)+76)) = int64(403726925889)
												*(*int32)(unsafe.Add(mBase, uint32(v6)+72)) = int32(1222)
												*(*int32)(unsafe.Add(mBase, uint32(v6)+68)) = int32(1223)
												*(*int32)(unsafe.Add(mBase, uint32(v6)+64)) = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v6)+60)) = int32(_a_F_PredicateLockShmemRequest_16)
												*(*int64)(unsafe.Add(mBase, uint32(v6)+52)) = int64(21474836480)
												v188 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockShmemRequest[4]))
												*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = v188
												F_SimpleLruRequestWithOpts(m, v6+int32(24))
												mBase = m.M
												v193 = m.ExcPending
												if v193 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = int32(_a_F_PredicateLockShmemRequest_17)
													*(*int64)(unsafe.Add(mBase, uint32(v6)+12)) = int64(16)
													*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = int32(_a_F_PredicateLockShmemRequest_18)
													F_ShmemRequestStructWithOpts(m, v6+int32(8))
													mBase = m.M
													v203 = m.ExcPending
													if v203 != 0 {
														return
													} else {
														m.G0 = v6 + int32(400)
														return
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
func F_PredicateLockTID(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v154 int32
	_ = v154
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockTID[0]))
	if v12 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+108)))
	if v16&int32(128) != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	F_ReleasePredicateLocks(m, int32(0), int32(1))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if base.Ui32(v23) < base.Ui32(int32(_a_F_PredicateLockTID_0)) {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	return
L8:
	;
	goto L1
L9:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+118)))
	if v27 == int32(116) {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	if v30 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	if base.Ui32(l3) < base.Ui32(int32(3)) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v166 = v23
	goto L13
L13:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = int64(4294967295)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v166
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v167
	v173 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockTID[1]))
	v174 = int32(0)
	v176 = F_hash_search(m, v173, v9, v174, v174)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L7
	} else {
		goto L55
	}
L14:
	;
	if v164 != 0 {
		goto L1
	} else {
		goto L54
	}
L15:
	;
	v164 = int32(0)
	goto L14
L16:
	;
	goto L17
L17:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockTID[2]))
	if v44 == l3 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v164 = int32(1)
	goto L14
L19:
	;
	goto L20
L20:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockTID[3]))
	if v48 <= int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v164 = v154
	goto L14
L22:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockTID[4]))
	if v52 == int32(0) {
		v154 = int32(0)
		goto L21
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockTID[5]))
	v124 = int32(0)
	v127 = v48 - int32(1)
	goto L44
L25:
	;
	v57 = v52
	goto L26
L26:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v57)+20))
	if v63 == int32(4) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v154 = int32(0)
	goto L21
L28:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v57)+80))
	if v117 != 0 {
		v57 = v117
		goto L26
	} else {
		goto L43
	}
L29:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	if v66 == int32(0) {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v69 = int32(1)
	if l3 == v66 {
		v154 = v69
		goto L21
	} else {
		goto L31
	}
L31:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v57)+52))
	v73 = v71 - int32(1)
	if v73 < int32(0) {
		goto L28
	} else {
		goto L32
	}
L32:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v57)+48))
	v79 = int32(0)
	v82 = v73
	goto L33
L33:
	;
	v87 = int32(2)
	v88 = base.I32_div_s(v82-v79, v87)
	v89 = v88 + v79
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v76+v89<<(uint(v87)%32))))
	if v93 == l3 {
		v154 = v69
		goto L21
	} else {
		goto L35
	}
L34:
	;
	goto L28
L35:
	;
	v102 = base.B2i32(v93-l3 < int32(0)) | base.B2i32(base.Ui32(v93) < base.Ui32(int32(3)))
	if v102 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v103 = v89 + int32(1)
	goto L38
L37:
	;
	v103 = v79
	goto L38
L38:
	;
	if v102 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v106 = v82
	goto L41
L40:
	;
	v106 = v89 - int32(1)
	goto L41
L41:
	;
	if v103 <= v106 {
		v79 = v103
		v82 = v106
		goto L33
	} else {
		goto L42
	}
L42:
	;
	goto L34
L43:
	;
	goto L27
L44:
	;
	v132 = int32(2)
	v133 = base.I32_div_s(v127-v124, v132)
	v134 = v133 + v124
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v122+v134<<(uint(v132)%32))))
	v139 = base.B2i32(v138 == l3)
	if v138 == l3 {
		v154 = v139
		goto L21
	} else {
		goto L46
	}
L45:
	;
	v154 = v139
	goto L21
L46:
	;
	v142 = base.B2i32(base.Ui32(v138) < base.Ui32(l3))
	if base.Ui32(v138) < base.Ui32(l3) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v143 = v134 + int32(1)
	goto L49
L48:
	;
	v143 = v124
	goto L49
L49:
	;
	if base.Ui32(v138) < base.Ui32(l3) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v146 = v127
	goto L52
L51:
	;
	v146 = v134 - int32(1)
	goto L52
L52:
	;
	if v143 <= v146 {
		v124 = v143
		v127 = v146
		goto L44
	} else {
		goto L53
	}
L53:
	;
	goto L45
L54:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v166 = v165
	goto L13
L55:
	;
	if v176 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176)+16)))
	if v178 != 0 {
		goto L1
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v179
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v181
	v183 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v183 | v184<<(uint(int32(16))%32)
	v189 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v189
	F_PredicateLockAcquire(m, v9)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L7
	} else {
		goto L60
	}
L59:
	;
	goto L58
L60:
	;
	goto L1
}
