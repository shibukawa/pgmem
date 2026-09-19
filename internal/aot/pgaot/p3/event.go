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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	v2 = l1
	v3 = l2
	v8 = m.G0
	v10 = v8 + int32(-64)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_EventTriggerSQLDropAddObject[0]))
	if v13 == int32(0) {
		m.G0 = v10 - int32(-64)
		return
	} else {
		v16 = int32(_a_F_EventTriggerSQLDropAddObject_0)
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_EventTriggerSQLDropAddObject[1]))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		*(*int32)(unsafe.Add(mBase, _c_F_EventTriggerSQLDropAddObject[1])) = v19
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
				v53 = v8 + int32(-48)
				v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				F_GetAttrDefaultColumnAddress(m, v53, v54)
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return
				} else {
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
					if v57 == int32(0) {
						v179 = F_getObjectIdentityParts(m, v22, v22+int32(28), v22+int32(32), int32(0))
						mBase = m.M
						v180 = m.ExcPending
						if v180 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v179
							v183 = F_getObjectTypeDescription(m, v22, int32(0))
							mBase = m.M
							v184 = m.ExcPending
							if v184 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v183
								v187 = *(*int32)(unsafe.Add(mBase, _c_F_EventTriggerSQLDropAddObject[0]))
								v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v188
								*(*int32)(unsafe.Add(mBase, uint32(v187)+4)) = v22 + int32(40)
								*(*int32)(unsafe.Add(mBase, _c_F_EventTriggerSQLDropAddObject[1])) = v17
								m.G0 = v10 - int32(-64)
								return
							}
						}
					} else {
						v60 = F_obtain_object_name_namespace(m, v53, v22)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return
						} else {
							if v60 != 0 {
								v179 = F_getObjectIdentityParts(m, v22, v22+int32(28), v22+int32(32), int32(0))
								mBase = m.M
								v180 = m.ExcPending
								if v180 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v179
									v183 = F_getObjectTypeDescription(m, v22, int32(0))
									mBase = m.M
									v184 = m.ExcPending
									if v184 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v183
										v187 = *(*int32)(unsafe.Add(mBase, _c_F_EventTriggerSQLDropAddObject[0]))
										v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v188
										*(*int32)(unsafe.Add(mBase, uint32(v187)+4)) = v22 + int32(40)
										*(*int32)(unsafe.Add(mBase, _c_F_EventTriggerSQLDropAddObject[1])) = v17
										m.G0 = v10 - int32(-64)
										return
									}
								}
							} else {
								F_pfree(m, v22)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_EventTriggerSQLDropAddObject[1])) = v17
									m.G0 = v10 - int32(-64)
									return
								}
							}
						}
					}
				}
			case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 13, 14, 15:
				v114 = F_obtain_object_name_namespace(m, l0, v22)
				mBase = m.M
				v115 = m.ExcPending
				if v115 != 0 {
					return
				} else {
					if v114 != 0 {
						v179 = F_getObjectIdentityParts(m, v22, v22+int32(28), v22+int32(32), int32(0))
						mBase = m.M
						v180 = m.ExcPending
						if v180 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v179
							v183 = F_getObjectTypeDescription(m, v22, int32(0))
							mBase = m.M
							v184 = m.ExcPending
							if v184 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v183
								v187 = *(*int32)(unsafe.Add(mBase, _c_F_EventTriggerSQLDropAddObject[0]))
								v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v188
								*(*int32)(unsafe.Add(mBase, uint32(v187)+4)) = v22 + int32(40)
								*(*int32)(unsafe.Add(mBase, _c_F_EventTriggerSQLDropAddObject[1])) = v17
								m.G0 = v10 - int32(-64)
								return
							}
						}
					} else {
						F_pfree(m, v22)
						mBase = m.M
						v117 = m.ExcPending
						if v117 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_EventTriggerSQLDropAddObject[1])) = v17
							m.G0 = v10 - int32(-64)
							return
						}
					}
				}
			case 11:
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v36 = *(*int32)(unsafe.Add(mBase, _c_F_EventTriggerSQLDropAddObject[2]))
				if base.B2i32(v36 != int32(0))&base.B2i32(v33 == v36) != 0 {
					v41 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v22)+38)) = uint8(v41)
					v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v167 = F_get_namespace_name(m, v166)
					mBase = m.M
					v168 = m.ExcPending
					if v168 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v167
						v179 = F_getObjectIdentityParts(m, v22, v22+int32(28), v22+int32(32), int32(0))
						mBase = m.M
						v180 = m.ExcPending
						if v180 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v179
							v183 = F_getObjectTypeDescription(m, v22, int32(0))
							mBase = m.M
							v184 = m.ExcPending
							if v184 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v183
								v187 = *(*int32)(unsafe.Add(mBase, _c_F_EventTriggerSQLDropAddObject[0]))
								v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v188
								*(*int32)(unsafe.Add(mBase, uint32(v187)+4)) = v22 + int32(40)
								*(*int32)(unsafe.Add(mBase, _c_F_EventTriggerSQLDropAddObject[1])) = v17
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
							v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v167 = F_get_namespace_name(m, v166)
							mBase = m.M
							v168 = m.ExcPending
							if v168 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v167
								v179 = F_getObjectIdentityParts(m, v22, v22+int32(28), v22+int32(32), int32(0))
								mBase = m.M
								v180 = m.ExcPending
								if v180 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v179
									v183 = F_getObjectTypeDescription(m, v22, int32(0))
									mBase = m.M
									v184 = m.ExcPending
									if v184 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v183
										v187 = *(*int32)(unsafe.Add(mBase, _c_F_EventTriggerSQLDropAddObject[0]))
										v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v188
										*(*int32)(unsafe.Add(mBase, uint32(v187)+4)) = v22 + int32(40)
										*(*int32)(unsafe.Add(mBase, _c_F_EventTriggerSQLDropAddObject[1])) = v17
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
								*(*int32)(unsafe.Add(mBase, _c_F_EventTriggerSQLDropAddObject[1])) = v17
								m.G0 = v10 - int32(-64)
								return
							}
						}
					}
				}
			case 16:
				v68 = F_table_open(m, int32(2620), int32(1))
				mBase = m.M
				v69 = m.ExcPending
				if v69 != 0 {
					return
				} else {
					v71 = v8 + int32(-48)
					v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					F_ScanKeyInit(m, v71, int32(1), int32(3), int32(184), v75)
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return
					} else {
						v78 = int32(0)
						v80 = int32(1)
						v83 = F_systable_beginscan(m, v68, int32(2702), v80, v78, v80, v71)
						mBase = m.M
						v84 = m.ExcPending
						if v84 != 0 {
							return
						} else {
							v85 = F_systable_getnext(m, v83)
							mBase = m.M
							v86 = m.ExcPending
							if v86 != 0 {
								return
							} else {
								if v85 != 0 {
									v87 = *(*int32)(unsafe.Add(mBase, uint32(v85)+16))
									v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+22)))
									v90 = *(*int32)(unsafe.Add(mBase, uint32(v87+v88)+4))
									v91 = v90
								} else {
									v91 = v78
								}
								F_systable_endscan(m, v83)
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
									return
								} else {
									F_relation_close(m, v68, int32(1))
									mBase = m.M
									v96 = m.ExcPending
									if v96 != 0 {
										return
									} else {
										if v91 == int32(0) {
											v179 = F_getObjectIdentityParts(m, v22, v22+int32(28), v22+int32(32), int32(0))
											mBase = m.M
											v180 = m.ExcPending
											if v180 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v179
												v183 = F_getObjectTypeDescription(m, v22, int32(0))
												mBase = m.M
												v184 = m.ExcPending
												if v184 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v183
													v187 = *(*int32)(unsafe.Add(mBase, _c_F_EventTriggerSQLDropAddObject[0]))
													v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v188
													*(*int32)(unsafe.Add(mBase, uint32(v187)+4)) = v22 + int32(40)
													*(*int32)(unsafe.Add(mBase, _c_F_EventTriggerSQLDropAddObject[1])) = v17
													m.G0 = v10 - int32(-64)
													return
												}
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v91
											*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(1259)
											v106 = F_obtain_object_name_namespace(m, v8+int32(-60), v22)
											mBase = m.M
											v107 = m.ExcPending
											if v107 != 0 {
												return
											} else {
												if v106 != 0 {
													v179 = F_getObjectIdentityParts(m, v22, v22+int32(28), v22+int32(32), int32(0))
													mBase = m.M
													v180 = m.ExcPending
													if v180 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v179
														v183 = F_getObjectTypeDescription(m, v22, int32(0))
														mBase = m.M
														v184 = m.ExcPending
														if v184 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v183
															v187 = *(*int32)(unsafe.Add(mBase, _c_F_EventTriggerSQLDropAddObject[0]))
															v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
															*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v188
															*(*int32)(unsafe.Add(mBase, uint32(v187)+4)) = v22 + int32(40)
															*(*int32)(unsafe.Add(mBase, _c_F_EventTriggerSQLDropAddObject[1])) = v17
															m.G0 = v10 - int32(-64)
															return
														}
													}
												} else {
													F_pfree(m, v22)
													mBase = m.M
													v109 = m.ExcPending
													if v109 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, _c_F_EventTriggerSQLDropAddObject[1])) = v17
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
					v122 = F_table_open(m, int32(3256), int32(1))
					mBase = m.M
					v123 = m.ExcPending
					if v123 != 0 {
						return
					} else {
						v125 = v8 + int32(-48)
						v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						F_ScanKeyInit(m, v125, int32(1), int32(3), int32(184), v129)
						mBase = m.M
						v131 = m.ExcPending
						if v131 != 0 {
							return
						} else {
							v132 = int32(0)
							v134 = int32(1)
							v137 = F_systable_beginscan(m, v122, int32(3257), v134, v132, v134, v125)
							mBase = m.M
							v138 = m.ExcPending
							if v138 != 0 {
								return
							} else {
								v139 = F_systable_getnext(m, v137)
								mBase = m.M
								v140 = m.ExcPending
								if v140 != 0 {
									return
								} else {
									if v139 != 0 {
										v141 = *(*int32)(unsafe.Add(mBase, uint32(v139)+16))
										v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+22)))
										v144 = *(*int32)(unsafe.Add(mBase, uint32(v141+v142)+68))
										v145 = v144
									} else {
										v145 = v132
									}
									F_systable_endscan(m, v137)
									mBase = m.M
									v147 = m.ExcPending
									if v147 != 0 {
										return
									} else {
										F_relation_close(m, v122, int32(1))
										mBase = m.M
										v150 = m.ExcPending
										if v150 != 0 {
											return
										} else {
											if v145 == int32(0) {
												v179 = F_getObjectIdentityParts(m, v22, v22+int32(28), v22+int32(32), int32(0))
												mBase = m.M
												v180 = m.ExcPending
												if v180 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v179
													v183 = F_getObjectTypeDescription(m, v22, int32(0))
													mBase = m.M
													v184 = m.ExcPending
													if v184 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v183
														v187 = *(*int32)(unsafe.Add(mBase, _c_F_EventTriggerSQLDropAddObject[0]))
														v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
														*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v188
														*(*int32)(unsafe.Add(mBase, uint32(v187)+4)) = v22 + int32(40)
														*(*int32)(unsafe.Add(mBase, _c_F_EventTriggerSQLDropAddObject[1])) = v17
														m.G0 = v10 - int32(-64)
														return
													}
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = int32(1)
												*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v145
												*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(1259)
												v160 = F_obtain_object_name_namespace(m, v8+int32(-60), v22)
												mBase = m.M
												v161 = m.ExcPending
												if v161 != 0 {
													return
												} else {
													if v160 != 0 {
														v179 = F_getObjectIdentityParts(m, v22, v22+int32(28), v22+int32(32), int32(0))
														mBase = m.M
														v180 = m.ExcPending
														if v180 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v179
															v183 = F_getObjectTypeDescription(m, v22, int32(0))
															mBase = m.M
															v184 = m.ExcPending
															if v184 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v183
																v187 = *(*int32)(unsafe.Add(mBase, _c_F_EventTriggerSQLDropAddObject[0]))
																v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
																*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v188
																*(*int32)(unsafe.Add(mBase, uint32(v187)+4)) = v22 + int32(40)
																*(*int32)(unsafe.Add(mBase, _c_F_EventTriggerSQLDropAddObject[1])) = v17
																m.G0 = v10 - int32(-64)
																return
															}
														}
													} else {
														F_pfree(m, v22)
														mBase = m.M
														v163 = m.ExcPending
														if v163 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, _c_F_EventTriggerSQLDropAddObject[1])) = v17
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
					v114 = F_obtain_object_name_namespace(m, l0, v22)
					mBase = m.M
					v115 = m.ExcPending
					if v115 != 0 {
						return
					} else {
						if v114 != 0 {
							v179 = F_getObjectIdentityParts(m, v22, v22+int32(28), v22+int32(32), int32(0))
							mBase = m.M
							v180 = m.ExcPending
							if v180 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v179
								v183 = F_getObjectTypeDescription(m, v22, int32(0))
								mBase = m.M
								v184 = m.ExcPending
								if v184 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v183
									v187 = *(*int32)(unsafe.Add(mBase, _c_F_EventTriggerSQLDropAddObject[0]))
									v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v188
									*(*int32)(unsafe.Add(mBase, uint32(v187)+4)) = v22 + int32(40)
									*(*int32)(unsafe.Add(mBase, _c_F_EventTriggerSQLDropAddObject[1])) = v17
									m.G0 = v10 - int32(-64)
									return
								}
							}
						} else {
							F_pfree(m, v22)
							mBase = m.M
							v117 = m.ExcPending
							if v117 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_EventTriggerSQLDropAddObject[1])) = v17
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
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
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
	return v88
L2:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v40 != 0 {
		goto L17
	} else {
		goto L18
	}
L3:
	;
	if v6 == int32(0) {
		v88 = v3
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
		goto L15
	}
L6:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if base.B2i32(v12 == int32(0))|base.B2i32(v12 != v15) != 0 {
		v33 = v12
		v34 = v15
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v33-v34 != 0 {
		v88 = v3
		goto L1
	} else {
		goto L14
	}
L8:
	;
	goto L7
L9:
	;
	v18 = v7
	v19 = v6
	goto L10
L10:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
	if v23 == int32(0) {
		v33 = v23
		v34 = v22
		goto L8
	} else {
		goto L12
	}
L11:
	;
	v33 = v23
	v34 = v22
	goto L8
L12:
	;
	v26 = int32(1)
	if v23 == v22 {
		v18 = v18 + v26
		v19 = v19 + v26
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	goto L2
L15:
	;
	return int32(0)
L16:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v74 = F_equal(m, v72, v73)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L30
	} else {
		goto L31
	}
L17:
	;
	if v39 == int32(0) {
		v88 = v3
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	if v39 == v40 {
		goto L16
	} else {
		goto L29
	}
L20:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
	if base.B2i32(v45 == int32(0))|base.B2i32(v45 != v48) != 0 {
		v66 = v45
		v67 = v48
		goto L22
	} else {
		goto L23
	}
L21:
	;
	if v66-v67 != 0 {
		v88 = v3
		goto L1
	} else {
		goto L28
	}
L22:
	;
	goto L21
L23:
	;
	v51 = v40
	v52 = v39
	goto L24
L24:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+1)))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+1)))
	if v56 == int32(0) {
		v66 = v56
		v67 = v55
		goto L22
	} else {
		goto L26
	}
L25:
	;
	v66 = v56
	v67 = v55
	goto L22
L26:
	;
	v59 = int32(1)
	if v56 == v55 {
		v51 = v51 + v59
		v52 = v52 + v59
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	goto L16
L29:
	;
	return int32(0)
L30:
	;
	return int32(0)
L31:
	;
	if v74 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	return int32(0)
L33:
	;
	goto L34
L34:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v84 = F_equal(m, v82, v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L30
	} else {
		goto L35
	}
L35:
	;
	v88 = v84
	goto L1
}
func F_event_trigger_in(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13854(m, l0, int32(_a_F_event_trigger_in_0), int32(367), int32(_a_F_event_trigger_in_1), int32(_a_F_event_trigger_in_2), int32(_a_F_event_trigger_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
