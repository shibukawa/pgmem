package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_EventTriggerSQLDropAddObject(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
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
	var v26 int64
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	v2 = l1
	v3 = l2
	v8 = m.G0
	v10 = v8 + int32(-64)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _consts[430]))
	if v13 == int32(0) {
		m.G0 = v10 - int32(-64)
		return
	} else {
		v16 = int32(4549024)
		v17 = *(*int32)(unsafe.Add(mBase, _consts[0]))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		*(*int32)(unsafe.Add(mBase, _consts[0])) = v19
		v22 = F_palloc0(m, int32(44))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v24
			v26 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
			*(*int64)(unsafe.Add(mBase, uint32(v22))) = v26
			*(*uint8)(unsafe.Add(mBase, uint32(v22)+37)) = uint8(v3)
			*(*uint8)(unsafe.Add(mBase, uint32(v22)+36)) = uint8(v2)
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			switch v30 - int32(2604) {
			case 0:
				v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				F_GetAttrDefaultColumnAddress(m, v8+int32(-48), v54)
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return
				} else {
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
					if v57 == int32(0) {
						v185 = F_getObjectIdentityParts(m, v22, v22+int32(28), v22+int32(32), int32(0))
						mBase = m.M
						v186 = m.ExcPending
						if v186 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v185
							v189 = F_getObjectTypeDescription(m, v22, int32(0))
							mBase = m.M
							v190 = m.ExcPending
							if v190 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v189
								v193 = *(*int32)(unsafe.Add(mBase, _consts[430]))
								v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v194
								*(*int32)(unsafe.Add(mBase, uint32(v193)+4)) = v22 + int32(40)
								*(*int32)(unsafe.Add(mBase, _consts[0])) = v17
								m.G0 = v10 - int32(-64)
								return
							}
						}
					} else {
						v62 = F_obtain_object_name_namespace(m, v8+int32(-48), v22)
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return
						} else {
							if v62 != 0 {
								v185 = F_getObjectIdentityParts(m, v22, v22+int32(28), v22+int32(32), int32(0))
								mBase = m.M
								v186 = m.ExcPending
								if v186 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v185
									v189 = F_getObjectTypeDescription(m, v22, int32(0))
									mBase = m.M
									v190 = m.ExcPending
									if v190 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v189
										v193 = *(*int32)(unsafe.Add(mBase, _consts[430]))
										v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v194
										*(*int32)(unsafe.Add(mBase, uint32(v193)+4)) = v22 + int32(40)
										*(*int32)(unsafe.Add(mBase, _consts[0])) = v17
										m.G0 = v10 - int32(-64)
										return
									}
								}
							} else {
								F_pfree(m, v22)
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, _consts[0])) = v17
									m.G0 = v10 - int32(-64)
									return
								}
							}
						}
					}
				}
			case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 13, 14, 15:
				v118 = F_obtain_object_name_namespace(m, l0, v22)
				mBase = m.M
				v119 = m.ExcPending
				if v119 != 0 {
					return
				} else {
					if v118 != 0 {
						v185 = F_getObjectIdentityParts(m, v22, v22+int32(28), v22+int32(32), int32(0))
						mBase = m.M
						v186 = m.ExcPending
						if v186 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v185
							v189 = F_getObjectTypeDescription(m, v22, int32(0))
							mBase = m.M
							v190 = m.ExcPending
							if v190 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v189
								v193 = *(*int32)(unsafe.Add(mBase, _consts[430]))
								v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v194
								*(*int32)(unsafe.Add(mBase, uint32(v193)+4)) = v22 + int32(40)
								*(*int32)(unsafe.Add(mBase, _consts[0])) = v17
								m.G0 = v10 - int32(-64)
								return
							}
						}
					} else {
						F_pfree(m, v22)
						mBase = m.M
						v121 = m.ExcPending
						if v121 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[0])) = v17
							m.G0 = v10 - int32(-64)
							return
						}
					}
				}
			case 11:
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v36 = *(*int32)(unsafe.Add(mBase, _consts[137]))
				if base.B2i32(v36 != int32(0))&base.B2i32(v33 == v36) != 0 {
					v41 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v22)+38)) = uint8(v41)
					v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v173 = F_get_namespace_name(m, v172)
					mBase = m.M
					v174 = m.ExcPending
					if v174 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v173
						v185 = F_getObjectIdentityParts(m, v22, v22+int32(28), v22+int32(32), int32(0))
						mBase = m.M
						v186 = m.ExcPending
						if v186 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v185
							v189 = F_getObjectTypeDescription(m, v22, int32(0))
							mBase = m.M
							v190 = m.ExcPending
							if v190 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v189
								v193 = *(*int32)(unsafe.Add(mBase, _consts[430]))
								v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v194
								*(*int32)(unsafe.Add(mBase, uint32(v193)+4)) = v22 + int32(40)
								*(*int32)(unsafe.Add(mBase, _consts[0])) = v17
								m.G0 = v10 - int32(-64)
								return
							}
						}
					}
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v44 = F_isAnyTempNamespace(m, v43)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						if v44 == int32(0) {
							v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v173 = F_get_namespace_name(m, v172)
							mBase = m.M
							v174 = m.ExcPending
							if v174 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v173
								v185 = F_getObjectIdentityParts(m, v22, v22+int32(28), v22+int32(32), int32(0))
								mBase = m.M
								v186 = m.ExcPending
								if v186 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v185
									v189 = F_getObjectTypeDescription(m, v22, int32(0))
									mBase = m.M
									v190 = m.ExcPending
									if v190 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v189
										v193 = *(*int32)(unsafe.Add(mBase, _consts[430]))
										v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v194
										*(*int32)(unsafe.Add(mBase, uint32(v193)+4)) = v22 + int32(40)
										*(*int32)(unsafe.Add(mBase, _consts[0])) = v17
										m.G0 = v10 - int32(-64)
										return
									}
								}
							}
						} else {
							F_pfree(m, v22)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[0])) = v17
								m.G0 = v10 - int32(-64)
								return
							}
						}
					}
				}
			case 16:
				v70 = F_table_open(m, int32(2620), int32(1))
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return
				} else {
					v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					F_ScanKeyInit(m, v8+int32(-48), int32(1), int32(3), int32(184), v77)
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return
					} else {
						v80 = int32(0)
						v82 = int32(1)
						v87 = F_systable_beginscan(m, v70, int32(2702), v82, v80, v82, v8+int32(-48))
						mBase = m.M
						v88 = m.ExcPending
						if v88 != 0 {
							return
						} else {
							v89 = F_systable_getnext(m, v87)
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return
							} else {
								if v89 != 0 {
									v91 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
									v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+22)))
									v94 = *(*int32)(unsafe.Add(mBase, uint32(v91+v92)+4))
									v95 = v94
								} else {
									v95 = v80
								}
								F_systable_endscan(m, v87)
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
									return
								} else {
									F_sequence_close(m, v70, int32(1))
									mBase = m.M
									v100 = m.ExcPending
									if v100 != 0 {
										return
									} else {
										if v95 == int32(0) {
											v185 = F_getObjectIdentityParts(m, v22, v22+int32(28), v22+int32(32), int32(0))
											mBase = m.M
											v186 = m.ExcPending
											if v186 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v185
												v189 = F_getObjectTypeDescription(m, v22, int32(0))
												mBase = m.M
												v190 = m.ExcPending
												if v190 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v189
													v193 = *(*int32)(unsafe.Add(mBase, _consts[430]))
													v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v194
													*(*int32)(unsafe.Add(mBase, uint32(v193)+4)) = v22 + int32(40)
													*(*int32)(unsafe.Add(mBase, _consts[0])) = v17
													m.G0 = v10 - int32(-64)
													return
												}
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v95
											*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(1259)
											v110 = F_obtain_object_name_namespace(m, v8+int32(-60), v22)
											mBase = m.M
											v111 = m.ExcPending
											if v111 != 0 {
												return
											} else {
												if v110 != 0 {
													v185 = F_getObjectIdentityParts(m, v22, v22+int32(28), v22+int32(32), int32(0))
													mBase = m.M
													v186 = m.ExcPending
													if v186 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v185
														v189 = F_getObjectTypeDescription(m, v22, int32(0))
														mBase = m.M
														v190 = m.ExcPending
														if v190 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v189
															v193 = *(*int32)(unsafe.Add(mBase, _consts[430]))
															v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
															*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v194
															*(*int32)(unsafe.Add(mBase, uint32(v193)+4)) = v22 + int32(40)
															*(*int32)(unsafe.Add(mBase, _consts[0])) = v17
															m.G0 = v10 - int32(-64)
															return
														}
													}
												} else {
													F_pfree(m, v22)
													mBase = m.M
													v113 = m.ExcPending
													if v113 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, _consts[0])) = v17
														m.G0 = v10 - int32(-64)
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
			default:
				if v30 == int32(3256) {
					v126 = F_table_open(m, int32(3256), int32(1))
					mBase = m.M
					v127 = m.ExcPending
					if v127 != 0 {
						return
					} else {
						v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						F_ScanKeyInit(m, v8+int32(-48), int32(1), int32(3), int32(184), v133)
						mBase = m.M
						v135 = m.ExcPending
						if v135 != 0 {
							return
						} else {
							v136 = int32(0)
							v138 = int32(1)
							v143 = F_systable_beginscan(m, v126, int32(3257), v138, v136, v138, v8+int32(-48))
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
								return
							} else {
								v145 = F_systable_getnext(m, v143)
								mBase = m.M
								v146 = m.ExcPending
								if v146 != 0 {
									return
								} else {
									if v145 != 0 {
										v147 = *(*int32)(unsafe.Add(mBase, uint32(v145)+16))
										v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+22)))
										v150 = *(*int32)(unsafe.Add(mBase, uint32(v147+v148)+68))
										v151 = v150
									} else {
										v151 = v136
									}
									F_systable_endscan(m, v143)
									mBase = m.M
									v153 = m.ExcPending
									if v153 != 0 {
										return
									} else {
										F_sequence_close(m, v126, int32(1))
										mBase = m.M
										v156 = m.ExcPending
										if v156 != 0 {
											return
										} else {
											if v151 == int32(0) {
												v185 = F_getObjectIdentityParts(m, v22, v22+int32(28), v22+int32(32), int32(0))
												mBase = m.M
												v186 = m.ExcPending
												if v186 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v185
													v189 = F_getObjectTypeDescription(m, v22, int32(0))
													mBase = m.M
													v190 = m.ExcPending
													if v190 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v189
														v193 = *(*int32)(unsafe.Add(mBase, _consts[430]))
														v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
														*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v194
														*(*int32)(unsafe.Add(mBase, uint32(v193)+4)) = v22 + int32(40)
														*(*int32)(unsafe.Add(mBase, _consts[0])) = v17
														m.G0 = v10 - int32(-64)
														return
													}
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = int32(1)
												*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v151
												*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(1259)
												v166 = F_obtain_object_name_namespace(m, v8+int32(-60), v22)
												mBase = m.M
												v167 = m.ExcPending
												if v167 != 0 {
													return
												} else {
													if v166 != 0 {
														v185 = F_getObjectIdentityParts(m, v22, v22+int32(28), v22+int32(32), int32(0))
														mBase = m.M
														v186 = m.ExcPending
														if v186 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v185
															v189 = F_getObjectTypeDescription(m, v22, int32(0))
															mBase = m.M
															v190 = m.ExcPending
															if v190 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v189
																v193 = *(*int32)(unsafe.Add(mBase, _consts[430]))
																v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
																*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v194
																*(*int32)(unsafe.Add(mBase, uint32(v193)+4)) = v22 + int32(40)
																*(*int32)(unsafe.Add(mBase, _consts[0])) = v17
																m.G0 = v10 - int32(-64)
																return
															}
														}
													} else {
														F_pfree(m, v22)
														mBase = m.M
														v169 = m.ExcPending
														if v169 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, _consts[0])) = v17
															m.G0 = v10 - int32(-64)
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
				} else {
					v118 = F_obtain_object_name_namespace(m, l0, v22)
					mBase = m.M
					v119 = m.ExcPending
					if v119 != 0 {
						return
					} else {
						if v118 != 0 {
							v185 = F_getObjectIdentityParts(m, v22, v22+int32(28), v22+int32(32), int32(0))
							mBase = m.M
							v186 = m.ExcPending
							if v186 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v185
								v189 = F_getObjectTypeDescription(m, v22, int32(0))
								mBase = m.M
								v190 = m.ExcPending
								if v190 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v189
									v193 = *(*int32)(unsafe.Add(mBase, _consts[430]))
									v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v194
									*(*int32)(unsafe.Add(mBase, uint32(v193)+4)) = v22 + int32(40)
									*(*int32)(unsafe.Add(mBase, _consts[0])) = v17
									m.G0 = v10 - int32(-64)
									return
								}
							}
						} else {
							F_pfree(m, v22)
							mBase = m.M
							v121 = m.ExcPending
							if v121 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[0])) = v17
								m.G0 = v10 - int32(-64)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F__equalCreateEventTrigStmt(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
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
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
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
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	v3 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v7 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v86
L2:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v39 != 0 {
		goto L18
	} else {
		goto L19
	}
L3:
	;
	if v6 == int32(0) {
		v86 = v3
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	if v6 == v7 {
		goto L2
	} else {
		goto L16
	}
L6:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v13 == int32(0) {
		v32 = v12
		v33 = v13
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v33-v32 != 0 {
		v86 = v3
		goto L1
	} else {
		goto L15
	}
L8:
	;
	goto L7
L9:
	;
	if v12 != v13 {
		v32 = v12
		v33 = v13
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v17 = v7
	v18 = v6
	goto L11
L11:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
	if v22 == int32(0) {
		v32 = v21
		v33 = v22
		goto L8
	} else {
		goto L13
	}
L12:
	;
	v32 = v21
	v33 = v22
	goto L8
L13:
	;
	v25 = int32(1)
	if v21 == v22 {
		v17 = v17 + v25
		v18 = v18 + v25
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	goto L2
L16:
	;
	return int32(0)
L17:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v72 = F_equal(m, v70, v71)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L32
	} else {
		goto L33
	}
L18:
	;
	if v38 == int32(0) {
		v86 = v3
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if v38 == v39 {
		goto L17
	} else {
		goto L31
	}
L21:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
	if v45 == int32(0) {
		v64 = v44
		v65 = v45
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if v65-v64 != 0 {
		v86 = v3
		goto L1
	} else {
		goto L30
	}
L23:
	;
	goto L22
L24:
	;
	if v44 != v45 {
		v64 = v44
		v65 = v45
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v49 = v39
	v50 = v38
	goto L26
L26:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+1)))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+1)))
	if v54 == int32(0) {
		v64 = v53
		v65 = v54
		goto L23
	} else {
		goto L28
	}
L27:
	;
	v64 = v53
	v65 = v54
	goto L23
L28:
	;
	v57 = int32(1)
	if v53 == v54 {
		v49 = v49 + v57
		v50 = v50 + v57
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	goto L17
L31:
	;
	return int32(0)
L32:
	;
	return int32(0)
L33:
	;
	if v72 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	return int32(0)
L35:
	;
	goto L36
L36:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v82 = F_equal(m, v80, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L32
	} else {
		goto L37
	}
L37:
	;
	v86 = v82
	goto L1
}
func F_event_trigger_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(233174)
			F_errmsg(m, int32(201439), v4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(513983), int32(367), int32(290468))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
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
